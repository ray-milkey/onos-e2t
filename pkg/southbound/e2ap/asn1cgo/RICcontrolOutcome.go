// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICcontrolOutcome.h"
import "C"
import (
	"encoding/binary"
	"unsafe"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newRicControlOutcome(rco *e2ap_commondatatypes.RiccontrolOutcome) *C.RICcontrolOutcome_t {
	return newOctetString(rco.GetValue())
}

func decodeRicControlOutcomeBytes(rcoBytes []byte) *e2ap_commondatatypes.RiccontrolOutcome {
	rcmC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(rcoBytes[:8])))),
		size: C.ulong(binary.LittleEndian.Uint64(rcoBytes[8:])),
	}
	return decodeRicControlOutcome(&rcmC)
}

func decodeRicControlOutcome(rcoC *C.RICcontrolOutcome_t) *e2ap_commondatatypes.RiccontrolOutcome {
	result := e2ap_commondatatypes.RiccontrolOutcome{
		Value: decodeOctetString(rcoC),
	}

	return &result
}
