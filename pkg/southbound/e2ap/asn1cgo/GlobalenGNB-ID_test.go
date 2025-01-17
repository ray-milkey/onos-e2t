// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createGlobalenGnbIDMsg() (*e2ap_ies.GlobalenGnbId, error) {

	globalenGnbID := e2ap_ies.GlobalenGnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: []byte{0x01, 0x02, 0x03},
		},
		GNbId: &e2ap_ies.EngnbId{
			EngnbId: &e2ap_ies.EngnbId_GNbId{
				GNbId: &asn1.BitString{
					Value: []byte{0xdc, 0xb8, 0x90, 0x00},
					Len:   32, //Should be of length 22 to 32
				},
			},
		},
	}

	//if err := globalenGnbID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating GlobalenGnbId %s", err.Error())
	//}
	return &globalenGnbID, nil
}

func Test_xerEncodingGlobalenGnbID(t *testing.T) {

	globalenGnbID, err := createGlobalenGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalenGnbId PDU")

	xer, err := xerEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 186, len(xer))
	t.Logf("GlobalenGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalenGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalenGnbID XER - decoded\n%v", result)
	assert.DeepEqual(t, globalenGnbID.GetPLmnIdentity().GetValue(), result.GetPLmnIdentity().GetValue())
	assert.Equal(t, globalenGnbID.GetGNbId().GetGNbId().GetLen(), result.GetGNbId().GetGNbId().GetLen())
	assert.DeepEqual(t, globalenGnbID.GetGNbId().GetGNbId().GetValue(), result.GetGNbId().GetGNbId().GetValue())
}

func Test_perEncodingGlobalenGnbID(t *testing.T) {

	globalenGnbID, err := createGlobalenGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalenGnbId PDU")

	per, err := perEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	//assert.Equal(t, 9, len(per))
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalenGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalenGnbID PER - decoded\n%v", result)
	assert.DeepEqual(t, globalenGnbID.GetPLmnIdentity().GetValue(), result.GetPLmnIdentity().GetValue())
	assert.Equal(t, globalenGnbID.GetGNbId().GetGNbId().GetLen(), result.GetGNbId().GetGNbId().GetLen())
	assert.DeepEqual(t, globalenGnbID.GetGNbId().GetGNbId().GetValue(), result.GetGNbId().GetGNbId().GetValue())
}
