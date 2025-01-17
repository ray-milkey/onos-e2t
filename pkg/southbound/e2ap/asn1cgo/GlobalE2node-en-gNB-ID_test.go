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

func createGlobalE2nodeEnGnbIDMsg() (*e2ap_ies.GlobalE2NodeEnGnbId, error) {

	globalE2nodeEnGnbID := e2ap_ies.GlobalE2NodeEnGnbId{
		GlobalEnGNbId: &e2ap_ies.GlobalenGnbId{
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
		},
		EnGNbCuUpId: &e2ap_ies.GnbCuUpId{
			Value: 123,
		},
		EnGNbDuId: &e2ap_ies.GnbDuId{
			Value: 2,
		},
	}

	//if err := globalE2nodeEnGnbID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating GlobalE2nodeEnGnbId %s", err.Error())
	//}
	return &globalE2nodeEnGnbID, nil
}

func Test_xerEncodingGlobalE2nodeEnGnbID(t *testing.T) {

	globalE2nodeEnGnbID, err := createGlobalE2nodeEnGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeEnGnbId PDU")

	xer, err := xerEncodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbID)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeEnGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalE2nodeEnGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeEnGnbID XER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeEnGnbID.GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, globalE2nodeEnGnbID.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), result.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, globalE2nodeEnGnbID.GetEnGNbCuUpId().GetValue(), result.GetEnGNbCuUpId().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetEnGNbDuId().GetValue(), result.GetEnGNbDuId().GetValue())
}

func Test_perEncodingGlobalE2nodeEnGnbID(t *testing.T) {

	globalE2nodeEnGnbID, err := createGlobalE2nodeEnGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeEnGnbId PDU")

	per, err := perEncodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbID)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeEnGnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalE2nodeEnGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeEnGnbID PER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeEnGnbID.GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, globalE2nodeEnGnbID.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), result.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, globalE2nodeEnGnbID.GetEnGNbCuUpId().GetValue(), result.GetEnGNbCuUpId().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetEnGNbDuId().GetValue(), result.GetEnGNbDuId().GetValue())
}
