// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func CreateRicServiceUpdateAcknowledgeE2apPdu(trID int32, rfAccepted types.RanFunctionRevisions) (*e2appdudescriptions.E2ApPdu, error) {

	if rfAccepted == nil {
		return nil, fmt.Errorf("RanFunctionsAccepetd was not passed - it is mandatory parameter")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						SuccessfulOutcome: &e2appducontents.RicserviceUpdateAcknowledge{
							ProtocolIes: &e2appducontents.RicserviceUpdateAcknowledgeIes{
								//E2ApProtocolIes9:  &ranFunctionsAccepted, //RAN functions Accepted
								//E2ApProtocolIes13: &ranFunctionsRejected, //RAN functions Rejected
								E2ApProtocolIes49: &e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes49{
									Id:          int32(v2.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2apies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
							Value: int32(v2.ProcedureCodeIDRICserviceUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	ranFunctionsAccepted := e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes9{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	for rfID, rfRevision := range rfAccepted {
		rfIDiIe := e2appducontents.RanfunctionIdItemIes{
			RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
				Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &e2appducontents.RanfunctionIdItem{
					RanFunctionId: &e2apies.RanfunctionId{
						Value: int32(rfID),
					},
					RanFunctionRevision: &e2apies.RanfunctionRevision{
						Value: int32(rfRevision),
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, &rfIDiIe)
	}
	e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetSuccessfulOutcome().GetProtocolIes().E2ApProtocolIes9 = &ranFunctionsAccepted

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
