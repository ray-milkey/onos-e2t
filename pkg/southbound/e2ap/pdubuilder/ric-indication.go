// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func RicIndicationE2apPdu(ricReqID types.RicRequest, ranFuncID types.RanFunctionID,
	ricAction e2apies.RicactionType, ricIndicationType e2apies.RicindicationType,
	ricIndHd types.RicIndicationHeader, ricIndMsg types.RicIndicationMessage) (
	*e2appdudescriptions.E2ApPdu, error) {

	ricRequestID := e2appducontents.RicindicationIes_RicindicationIes29{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionID := e2appducontents.RicindicationIes_RicindicationIes5{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: int32(ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricAct := e2appducontents.RicindicationIes_RicindicationIes15{
		Id:          int32(v2.ProtocolIeIDRicactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicactionId{
			Value: int32(ricAction),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricIndType := e2appducontents.RicindicationIes_RicindicationIes28{
		Id:          int32(v2.ProtocolIeIDRicindicationType),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ricIndicationType,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricIndHeader := e2appducontents.RicindicationIes_RicindicationIes25{
		Id:          int32(v2.ProtocolIeIDRicindicationHeader),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RicindicationHeader{
			Value: []byte(ricIndHd),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricIndMessage := e2appducontents.RicindicationIes_RicindicationIes26{
		Id:          int32(v2.ProtocolIeIDRicindicationMessage),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RicindicationMessage{
			Value: []byte(ricIndMsg),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicIndication: &e2appdudescriptions.RicIndication{
						InitiatingMessage: &e2appducontents.Ricindication{
							ProtocolIes: &e2appducontents.RicindicationIes{
								E2ApProtocolIes29: &ricRequestID,  // RIC Requestor & RIC Instance ID
								E2ApProtocolIes5:  &ranFunctionID, // RAN function ID
								E2ApProtocolIes15: &ricAct,        // RIC Action
								//E2ApProtocolIes27: &ricIndicationSn,  // RIC Indication Sn (Sequence Number?)
								E2ApProtocolIes28: &ricIndType,    // RIC Indication Type
								E2ApProtocolIes25: &ricIndHeader,  // RIC Indication Header
								E2ApProtocolIes26: &ricIndMessage, // RIC Indication Message
								//E2ApProtocolIes20: &ricCallProcessID, // RIC Call Process ID
							},
						},
						ProcedureCode: &e2ap_constants.IdRicindication{
							Value: int32(v2.ProcedureCodeIDRICindication),
						},
						Criticality: &e2ap_commondatatypes.CriticalityIgnore{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
						},
					},
				},
			},
		},
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
