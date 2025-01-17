// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2ap_ies/pdubuilder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createGlobalngeNbIDMsg() (*e2ap_ies.GlobalngeNbId, error) {

	globalngeNbID := e2ap_ies.GlobalngeNbId{
		PlmnId: &e2ap_commondatatypes.PlmnIdentity{
			Value: []byte{0x01, 0x02, 0x03},
		},
		EnbId: &e2ap_ies.EnbIdChoice{
			EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdMacro{
				EnbIdMacro: &asn1.BitString{
					Value: []byte{0xdc, 0xb8, 0x90},
					Len:   20,
				},
			},
		},
	}

	//if err := globalngeNbID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating GlobalngeNbId %s", err.Error())
	//}
	return &globalngeNbID, nil
}

func Test_xerEncodingGlobalngeNbID(t *testing.T) {

	globalngeNbID, err := createGlobalngeNbIDMsg()
	assert.NilError(t, err, "Error creating GlobalngeNbID PDU")
	//t.Logf("GlobalngeNbID is\n%v", globalngeNbID)

	xer, err := xerEncodeGlobalngeNbID(globalngeNbID)
	assert.NilError(t, err)
	t.Logf("GlobalngeNbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalngeNbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalngeNbID XER - decoded\n%v", result)
	assert.DeepEqual(t, globalngeNbID.GetPlmnId().GetValue(), result.GetPlmnId().GetValue())
	assert.DeepEqual(t, globalngeNbID.GetEnbId().GetEnbIdMacro().GetValue(), result.GetEnbId().GetEnbIdMacro().GetValue())
	assert.Equal(t, globalngeNbID.GetEnbId().GetEnbIdMacro().GetLen(), result.GetEnbId().GetEnbIdMacro().GetLen())
}

func Test_perEncodingGlobalngeNbID(t *testing.T) {

	globalngeNbID, err := createGlobalngeNbIDMsg()
	assert.NilError(t, err, "Error creating GlobalngeNbID PDU")

	per, err := perEncodeGlobalngeNbID(globalngeNbID)
	assert.NilError(t, err)
	t.Logf("GlobalngeNbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalngeNbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalngeNbID PER - decoded\n%v", result)
	assert.DeepEqual(t, globalngeNbID.GetPlmnId().GetValue(), result.GetPlmnId().GetValue())
	assert.DeepEqual(t, globalngeNbID.GetEnbId().GetEnbIdMacro().GetValue(), result.GetEnbId().GetEnbIdMacro().GetValue())
	assert.DeepEqual(t, globalngeNbID.GetEnbId().GetEnbIdMacro().GetLen(), result.GetEnbId().GetEnbIdMacro().GetLen())
}
