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

func createGlobalE2nodeNgEnbIDMsg() (*e2ap_ies.GlobalE2NodeNgEnbId, error) {
	globalE2nodeNgEnbID := e2ap_ies.GlobalE2NodeNgEnbId{
		GlobalNgENbId: &e2ap_ies.GlobalngeNbId{
			PlmnId: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			EnbId: &e2ap_ies.EnbIdChoice{
				EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdMacro{
					EnbIdMacro: &asn1.BitString{
						Value: []byte{0x4d, 0xcb, 0xb0},
						Len:   20,
					},
				},
			},
		},
		NgEnbDuId: &e2ap_ies.NgenbDuId{
			Value: 2,
		},
	}

	//if err := globalE2nodeNgEnbID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating GlobalE2nodeNgEnbId %s", err.Error())
	//}
	return &globalE2nodeNgEnbID, nil
}

func Test_xerEncodingGlobalE2nodeNgEnbID(t *testing.T) {

	globalE2nodeNgEnbID, err := createGlobalE2nodeNgEnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeNgEnbId PDU")
	t.Logf("GlobalE2nodeNgEnbID to be encoded\n%v", globalE2nodeNgEnbID)

	xer, err := xerEncodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbID)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeNgEnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalE2nodeNgEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeNgEnbID XER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetPlmnId().GetValue(), result.GetGlobalNgENbId().GetPlmnId().GetValue())
	assert.Equal(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue())
	assert.Equal(t, globalE2nodeNgEnbID.GetNgEnbDuId().GetValue(), result.GetNgEnbDuId().GetValue())
}

func Test_perEncodingGlobalE2nodeNgEnbID(t *testing.T) {

	globalE2nodeNgEnbID, err := createGlobalE2nodeNgEnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeNgEnbId PDU")

	per, err := perEncodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbID)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeNgEnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalE2nodeNgEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeNgEnbID PER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetPlmnId().GetValue(), result.GetGlobalNgENbId().GetPlmnId().GetValue())
	assert.Equal(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue())
	assert.Equal(t, globalE2nodeNgEnbID.GetNgEnbDuId().GetValue(), result.GetNgEnbDuId().GetValue())
}
