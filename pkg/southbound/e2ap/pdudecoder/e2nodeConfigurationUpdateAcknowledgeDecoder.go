// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"

	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2nodeConfigurationUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32,
	[]*types.E2NodeComponentConfigUpdateAckItem, error) {
	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	e2ncua := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate()
	if e2ncua == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdateAcknowledge")
	}

	e2nccual := make([]*types.E2NodeComponentConfigUpdateAckItem, 0)
	list := e2ncua.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()
	for _, ie := range list {
		e2nccuai := types.E2NodeComponentConfigUpdateAckItem{}
		e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentInterfaceType()
		e2nccuai.E2NodeComponentID = ie.GetValue().GetE2NodeComponentId()
		e2nccuai.E2NodeComponentConfigurationAck = types.E2NodeComponentConfigurationAck{
			UpdateOutcome: ie.GetValue().GetE2NodeComponentConfigurationAck().GetUpdateOutcome(),
			FailureCause:  ie.GetValue().GetE2NodeComponentConfigurationAck().GetFailureCause(),
		}

		e2nccual = append(e2nccual, &e2nccuai)
	}

	transactionID := e2ncua.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue()

	return &transactionID, e2nccual, nil
}
