// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "RICeventTriggerDefinition.h"
import "C"
import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newRicEventTriggerDefinition(retd *e2ap_commondatatypes.RiceventTriggerDefinition) *C.RICeventTriggerDefinition_t {
	return newOctetString(retd.GetValue())
}

func decodeRicEventTriggerDefinition(retdC *C.RICeventTriggerDefinition_t) *e2ap_commondatatypes.RiceventTriggerDefinition {
	result := e2ap_commondatatypes.RiceventTriggerDefinition{
		Value: decodeOctetString(retdC),
	}

	return &result
}
