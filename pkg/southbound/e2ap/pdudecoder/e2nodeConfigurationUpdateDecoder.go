// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"

	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2nodeConfigurationUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, *types.E2NodeIdentity, []*types.E2NodeComponentConfigUpdateItem, error) {
	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	e2ncu := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate()
	if e2ncu == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdate")
	}

	transactionID := e2ncu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue()

	globalE2NodeID := e2ncu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes3().GetValue()
	nodeIdentity, err := ExtractE2NodeIdentity(globalE2NodeID)
	if err != nil {
		return nil, nil, nil, err
	}

	e2nccual := make([]*types.E2NodeComponentConfigUpdateItem, 0)
	list := e2ncu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()
	for _, ie := range list {
		e2nccuai := types.E2NodeComponentConfigUpdateItem{}
		e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentInterfaceType()
		e2nccuai.E2NodeComponentID = ie.GetValue().GetE2NodeComponentId()
		e2nccuai.E2NodeComponentConfiguration = *ie.GetValue().GetE2NodeComponentConfiguration()

		e2nccual = append(e2nccual, &e2nccuai)
	}

	return &transactionID, nodeIdentity, e2nccual, nil
}
