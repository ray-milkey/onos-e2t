// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func createE2nodeConfigurationUpdateFailureMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {

	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V2S
	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2nodeConfigurationUpdateFailure, err := pdubuilder.CreateE2NodeConfigurationUpdateFailureE2apPdu(1, &e2apies.Cause{
		Cause: &e2apies.Cause_Protocol{
			Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
		},
	})
	if err != nil {
		return nil, err
	}

	e2nodeConfigurationUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetUnsuccessfulOutcome().
		SetTimeToWait(ttw).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
			},
		})

	//if err := e2nodeConfigurationUpdateFailure.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeConfigurationUpdateFailure %s", err.Error())
	//}
	return e2nodeConfigurationUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetUnsuccessfulOutcome(), nil
}

func Test_xerEncodingE2nodeConfigurationUpdateFailure(t *testing.T) {

	e2nodeConfigurationUpdateFailure, err := createE2nodeConfigurationUpdateFailureMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateFailure PDU")

	xer, err := xerEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdateFailure XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdateFailure(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateFailure XER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}

func Test_perEncodingE2nodeConfigurationUpdateFailure(t *testing.T) {

	e2nodeConfigurationUpdateFailure, err := createE2nodeConfigurationUpdateFailureMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateFailure PDU")

	per, err := perEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdateFailure PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdateFailure(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateFailure PER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}
