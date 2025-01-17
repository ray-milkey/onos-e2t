// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"testing"

	"gotest.tools/assert"
)

const RicreqidXer = `<RICrequestID>
	<ricRequestorID>543210</ricRequestorID>
	<ricInstanceID>6789</ricInstanceID>
</RICrequestID>`

func Test_RicRequestID(t *testing.T) {
	result, err := xerDecodeRicRequestID([]byte(RicreqidXer))
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Equal(t, int32(543210), result.RicRequestorId)
	assert.Equal(t, int32(6789), result.RicInstanceId)
}
