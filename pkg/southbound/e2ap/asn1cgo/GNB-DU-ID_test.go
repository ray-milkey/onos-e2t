// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createGnbDuID() *e2apies.GnbDuId {

	return &e2apies.GnbDuId{
		Value: 1234,
	}
}

func Test_xerEncodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	xer, err := xerEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	assert.Equal(t, 28, len(xer))
	t.Logf("GnbDuID XER\n%s", string(xer))
}

func Test_xerDecodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	xer, err := xerEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	assert.Equal(t, 28, len(xer))
	t.Logf("GnbDuID XER\n%s", string(xer))

	result, err := xerDecodeGnbDuID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbDuID XER - decoded\n%s", result)
}

func Test_perEncodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	per, err := perEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))
}

func Test_perDecodeGnbDuID(t *testing.T) {

	gnbDuID := createGnbDuID()

	per, err := perEncodeGnbDuID(gnbDuID)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("GnbDuID PER\n%v", hex.Dump(per))

	result, err := perDecodeGnbDuID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GnbDuID PER - decoded\n%s", result)
}
