// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/encoding"

	"github.com/google/uuid"
	epapi "github.com/onosproject/onos-e2sub/api/e2/endpoint/v1beta1"
	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	subtaskapi "github.com/onosproject/onos-e2sub/api/e2/task/v1beta1"
	e2tapi "github.com/onosproject/onos-e2t/api/ricapi/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/app"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/connection"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/endpoint"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscriptiontask"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/termination"
)

var log = logging.GetLogger("e2")

const defaultServicePort = 5150

// Config is an E2 client configuration
type Config struct {
	// AppID is the application identifier
	AppID app.ID
	// InstanceID is the application instance identifier
	InstanceID app.InstanceID
	// SubscriptionService is the subscription service configuration
	SubscriptionService ServiceConfig
}

// ServiceConfig is an E2 service configuration
type ServiceConfig struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int
}

// GetHost gets the service host
func (c ServiceConfig) GetHost() string {
	return c.Host
}

// GetPort gets the service port
func (c ServiceConfig) GetPort() int {
	if c.Port == 0 {
		return defaultServicePort
	}
	return c.Port
}

// Client is an E2 client
type Client interface {
	// Subscribe subscribes the client to indications
	Subscribe(ctx context.Context, sub subscription.Subscription, ch chan<- indication.Indication) error
}

// NewClient creates a new E2 client
func NewClient(config Config) (Client, error) {
	uuid.SetNodeID([]byte(fmt.Sprintf("%s:%s", config.AppID, config.InstanceID)))
	conns := connection.NewManager()
	subConn, err := conns.Connect(fmt.Sprintf("%s:%d", config.SubscriptionService.GetHost(), config.SubscriptionService.GetPort()))
	if err != nil {
		return nil, err
	}
	return &e2Client{
		config:     config,
		epClient:   endpoint.NewClient(subConn),
		subClient:  subscription.NewClient(subConn),
		taskClient: subscriptiontask.NewClient(subConn),
		conns:      conns,
	}, nil
}

// e2Client is the default E2 client implementation
type e2Client struct {
	config     Config
	epClient   endpoint.Client
	subClient  subscription.Client
	taskClient subscriptiontask.Client
	conns      *connection.Manager
}

func getEncodingType(subscription subscription.Subscription) subapi.Encoding {
	var encodingType subapi.Encoding
	switch subscription.EncodingType.String() {
	case encoding.ASN1.String():
		encodingType = subapi.Encoding_ENCODING_ASN1
	case encoding.PROTO.String():
		encodingType = subapi.Encoding_ENCODING_PROTO
	default:
		encodingType = subapi.Encoding_ENCODING_PROTO

	}
	return encodingType

}

func getPayload(subscription subscription.Subscription) ([]byte, error) {
	var payload []byte
	var err error
	switch subscription.EncodingType.String() {
	case encoding.ASN1.String():
		// TODO encode the payload in ASN.1
	case encoding.PROTO.String():
		payload, err = subscription.Payload.GetProtoValue()

	default:
		payload, err = subscription.Payload.GetProtoValue()

	}
	return payload, err
}

func (c *e2Client) Subscribe(ctx context.Context, subscription subscription.Subscription, ch chan<- indication.Indication) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	encodingType := getEncodingType(subscription)
	bytes, err := getPayload(subscription)
	if err != nil {
		return err
	}

	sub := &subapi.Subscription{
		ID:       subapi.ID(id.String()),
		AppID:    subapi.AppID(c.config.AppID),
		E2NodeID: subapi.E2NodeID(subscription.NodeID),
		ServiceModel: &subapi.ServiceModel{
			ID: subapi.ServiceModelID(subscription.ServiceModel.ID),
		},
		Payload: &subapi.Payload{
			Encoding: encodingType,
			Bytes:    bytes,
		},
	}

	client := &subscriptionClient{
		e2Client: c,
		sub:      sub,
		ch:       ch,
		ctx:      ctx,
	}
	return client.subscribe()
}

// subscriptionClient is a client for managing a subscription
type subscriptionClient struct {
	*e2Client
	sub    *subapi.Subscription
	ch     chan<- indication.Indication
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *subscriptionClient) subscribe() error {
	err := c.subClient.Add(c.ctx, c.sub)
	if err != nil {
		return err
	}

	taskCh := make(chan subtaskapi.Event)
	err = c.taskClient.Watch(c.ctx, taskCh, subscriptiontask.WithSubscriptionID(c.sub.ID))
	if err != nil {
		return err
	}
	go c.processEvents(taskCh)
	return nil
}

func (c *subscriptionClient) processEvents(eventCh <-chan subtaskapi.Event) {
	for event := range eventCh {
		if event.Type == subtaskapi.EventType_NONE || event.Type == subtaskapi.EventType_CREATED {
			err := c.stream(event.Task.EndpointID)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (c *subscriptionClient) stream(epID epapi.ID) error {
	response, err := c.epClient.Get(c.ctx, epID)
	if err != nil {
		return err
	}

	conn, err := c.conns.Connect(fmt.Sprintf("%s:%d", response.IP, response.Port))
	if err != nil {
		return err
	}

	if c.cancel != nil {
		c.cancel()
	}

	client := termination.NewClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	responseCh := make(chan e2tapi.StreamResponse)
	requestCh, err := client.Stream(ctx, responseCh)
	if err != nil {
		cancel()
		return err
	}
	c.cancel = cancel

	requestCh <- e2tapi.StreamRequest{
		AppID:          e2tapi.AppID(c.config.AppID),
		InstanceID:     e2tapi.InstanceID(c.config.InstanceID),
		SubscriptionID: e2tapi.SubscriptionID(c.sub.ID),
	}

	go func() {
		for response := range responseCh {
			c.ch <- indication.Indication{
				EncodingType: encoding.Type(response.Header.EncodingType),
				Payload: indication.Payload{
					Value: response.Payload,
				},
			}
		}
	}()
	return nil
}
