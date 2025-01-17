// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createE2connectionUpdateItemMsg() (*e2ap_pdu_contents.E2ConnectionUpdateItem, error) {

	e2connectionUpdateItem := e2ap_pdu_contents.E2ConnectionUpdateItem{
		TnlInformation: &e2ap_ies.Tnlinformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xcd, 0x9b},
				Len:   16,
			},
			TnlAddress: &asn1.BitString{
				Value: []byte{0xab, 0xbc, 0xcd, 0xde, 0xef, 0xf5, 0xd6, 0xb7},
				Len:   64,
			},
		},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH,
	}

	//if err := e2connectionUpdateItem.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2connectionUpdateItem %s", err.Error())
	//}
	return &e2connectionUpdateItem, nil
}

func Test_xerEncodingE2connectionUpdateItem(t *testing.T) {

	e2connectionUpdateItem, err := createE2connectionUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateItem PDU")

	xer, err := xerEncodeE2connectionUpdateItem(e2connectionUpdateItem)
	assert.NilError(t, err)
	assert.Equal(t, 315, len(xer))
	t.Logf("E2connectionUpdateItem XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateItem XER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateItem.GetTnlInformation().GetTnlPort().GetValue(), result.GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateItem.GetTnlInformation().GetTnlAddress().GetValue(), result.GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdateItem.GetTnlUsage(), result.GetTnlUsage())
}

func Test_perEncodingE2connectionUpdateItem(t *testing.T) {

	e2connectionUpdateItem, err := createE2connectionUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateItem PDU")

	per, err := perEncodeE2connectionUpdateItem(e2connectionUpdateItem)
	assert.NilError(t, err)
	assert.Equal(t, 13, len(per))
	t.Logf("E2connectionUpdateItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateItem PER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateItem.GetTnlInformation().GetTnlPort().GetValue(), result.GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateItem.GetTnlInformation().GetTnlAddress().GetValue(), result.GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdateItem.GetTnlUsage(), result.GetTnlUsage())
}
