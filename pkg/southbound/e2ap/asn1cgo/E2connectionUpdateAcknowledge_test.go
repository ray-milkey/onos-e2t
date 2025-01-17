// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createE2connectionUpdateAcknowledgeMsg() (*e2ap_pdu_contents.E2ConnectionUpdateAcknowledge, error) {

	e2connectionUpdateAcknowledge, err := pdubuilder.CreateE2connectionUpdateAcknowledgeE2apPdu(1)
	if err != nil {
		return nil, err
	}

	e2connectionUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetSuccessfulOutcome().
		SetE2ConnectionSetup([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}}).SetE2ConnectionSetupFailed([]*types.E2ConnectionSetupFailedItem{{TnlInformation: types.TnlInformation{
		TnlPort: &asn1.BitString{
			Value: []byte{0xae, 0x89},
			Len:   16,
		},
		TnlAddress: asn1.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
			Len:   64,
		}},
		Cause: e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_Protocol{
				Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
			}}}})

	//if err := e2connectionUpdateAcknowledge.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2connectionUpdateAcknowledge %s", err.Error())
	//}
	return e2connectionUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetSuccessfulOutcome(), nil
}

func Test_xerEncodingE2connectionUpdateAcknowledge(t *testing.T) {

	e2connectionUpdateAcknowledge, err := createE2connectionUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateAcknowledge PDU")

	xer, err := xerEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateAcknowledge(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateAcknowledge XER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetCause().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetCause().GetProtocol())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}

func Test_perEncodingE2connectionUpdateAcknowledge(t *testing.T) {

	e2connectionUpdateAcknowledge, err := createE2connectionUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateAcknowledge PDU")

	per, err := perEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateAcknowledge(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateAcknowledge PER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes39().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetCause().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes40().GetValue().GetValue()[0].GetValue().GetCause().GetProtocol())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}
