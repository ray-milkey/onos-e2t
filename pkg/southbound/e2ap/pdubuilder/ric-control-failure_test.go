// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicControlFailure(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	//var ricCallPrID types.RicCallProcessID = []byte("123")
	//var ricCtrlOut types.RicControlOutcome = []byte("456")
	cause := e2apies.Cause{
		Cause: &e2apies.Cause_RicRequest{
			RicRequest: e2apies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN,
		},
	}
	newE2apPdu, err := CreateRicControlFailureE2apPdu(ricRequestID,
		ranFuncID, &cause)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome().
		SetRicControlOutcome([]byte{0xFF, 0xFF, 0xDD, 0x4A}).SetRicCallProcessID([]byte{0xCC, 0x3D, 0x1F})

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RIC Control Request XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RIC Control Request E2AP PDU\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}
