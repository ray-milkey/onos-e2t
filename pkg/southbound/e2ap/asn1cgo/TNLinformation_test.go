// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createTnlinformationMsg() (*e2ap_ies.Tnlinformation, error) {

	bs1 := &asn1.BitString{
		Value: []byte{0xab, 0xbc, 0xcd, 0xde, 0xef, 0xf5, 0xd6, 0xb7},
		Len:   64,
	}

	bs2 := &asn1.BitString{
		Value: []byte{0xcd, 0x9b},
		Len:   16,
	}

	tnlinformation := e2ap_ies.Tnlinformation{
		TnlAddress: bs1,
		TnlPort:    bs2,
	}

	//if err := tnlinformation.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating Tnlinformation %s", err.Error())
	//}
	return &tnlinformation, nil
}

func Test_xerEncodingTnlinformation(t *testing.T) {

	tnlinformation, err := createTnlinformationMsg()
	assert.NilError(t, err, "Error creating TNLinformation PDU")
	t.Logf("TNLinformation (message)\n%v", tnlinformation)

	xer, err := xerEncodeTnlinformation(tnlinformation)
	assert.NilError(t, err)
	assert.Equal(t, 197, len(xer))
	t.Logf("TNLinformation XER\n%s", string(xer))

	result, err := xerDecodeTnlinformation(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLinformation XER - decoded\n%v", result)

	assert.DeepEqual(t, tnlinformation.GetTnlAddress().GetValue(), result.GetTnlAddress().GetValue())
	assert.Equal(t, tnlinformation.GetTnlAddress().GetLen(), result.GetTnlAddress().GetLen())
	assert.DeepEqual(t, tnlinformation.GetTnlPort().GetValue(), result.GetTnlPort().GetValue())
	assert.Equal(t, tnlinformation.GetTnlPort().GetLen(), result.GetTnlPort().GetLen())
}

func Test_perEncodingTnlinformation(t *testing.T) {

	tnlinformation, err := createTnlinformationMsg()
	assert.NilError(t, err, "Error creating TNLinformation PDU")
	t.Logf("TNLinformation (message)\n%v", tnlinformation)

	per, err := perEncodeTnlinformation(tnlinformation)
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("TNLinformation PER\n%v", hex.Dump(per))

	result, err := perDecodeTnlinformation(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLinformation PER - decoded\n%v", result)

	assert.DeepEqual(t, tnlinformation.GetTnlAddress().GetValue(), result.GetTnlAddress().GetValue())
	assert.Equal(t, tnlinformation.GetTnlAddress().GetLen(), result.GetTnlAddress().GetLen())
	assert.DeepEqual(t, tnlinformation.GetTnlPort().GetValue(), result.GetTnlPort().GetValue())
	assert.Equal(t, tnlinformation.GetTnlPort().GetLen(), result.GetTnlPort().GetLen())
}
