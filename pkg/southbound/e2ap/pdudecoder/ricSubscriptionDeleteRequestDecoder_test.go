// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"io/ioutil"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"gotest.tools/assert"
)

func Test_DecodeRicSubscriptionDeleteRequestPdu(t *testing.T) {
	ricSubscriptionDeleteRequestXer, err := ioutil.ReadFile("../test/RICsubscriptionDeleteRequest.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionDeleteRequestXer)
	assert.NilError(t, err)

	ricReq, ranFuncID, err := DecodeRicSubscriptionDeleteRequestPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(ricReq.RequestorID))
	assert.Equal(t, 2, int(ricReq.InstanceID))
	assert.Equal(t, 3, int(ranFuncID))
}
