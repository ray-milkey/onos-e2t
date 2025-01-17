// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICactionDefinition.h"
import "C"
import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newRicActionDefinition(rad *e2ap_commondatatypes.RicactionDefinition) *C.RICactionDefinition_t {
	return newOctetString(rad.Value)
}

func decodeRicActionDefinition(radC *C.RICactionDefinition_t) *e2ap_commondatatypes.RicactionDefinition {
	result := e2ap_commondatatypes.RicactionDefinition{
		Value: decodeOctetString(radC),
	}

	return &result
}
