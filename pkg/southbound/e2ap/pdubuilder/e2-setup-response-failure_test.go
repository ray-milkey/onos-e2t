// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"testing"

	"gotest.tools/assert"
)

func TestE2SetupResponseFailure(t *testing.T) {
	newE2apPdu, err := CreateSetupResponseFailureE2apPdu(21, 1, 10)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
}
