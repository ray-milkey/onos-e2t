// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2connectionUpdateFailurePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, *e2ap_ies.Cause, *e2ap_ies.TimeToWait,
	*v2.ProcedureCodeT, *e2ap_commondatatypes.Criticality, *e2ap_commondatatypes.TriggeringMessage, *types.RicRequest,
	[]*types.CritDiag, error) {
	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	e2cuf := e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate()
	if e2cuf == nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2connectionUpdateFailure")
	}

	cause := e2cuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes1().GetValue()
	ttw := e2cuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes31().GetValue()

	critDiagnostics := e2cuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2()
	var pc v2.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID types.RicRequest
	var diags []*types.CritDiag
	if critDiagnostics != nil { //It's optional
		pc = v2.ProcedureCodeT(critDiagnostics.GetValue().GetProcedureCode().GetValue())
		crit = critDiagnostics.GetValue().GetProcedureCriticality()
		tm = critDiagnostics.GetValue().GetTriggeringMessage()
		critDiagRequestID = types.RicRequest{
			RequestorID: types.RicRequestorID(critDiagnostics.GetValue().GetRicRequestorId().GetRicRequestorId()),
			InstanceID:  types.RicInstanceID(critDiagnostics.GetValue().GetRicRequestorId().GetRicInstanceId()),
		}
		for _, ie := range critDiagnostics.GetValue().GetIEsCriticalityDiagnostics().GetValue() {
			diag := types.CritDiag{
				IECriticality: ie.IEcriticality,
				IEId:          v2.ProtocolIeID(ie.GetIEId().GetValue()),
				TypeOfError:   ie.TypeOfError,
			}
			diags = append(diags, &diag)
		}
	}

	transactionID := e2cuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue()

	return &transactionID, cause, &ttw, &pc, &crit, &tm, &critDiagRequestID, diags, nil
}
