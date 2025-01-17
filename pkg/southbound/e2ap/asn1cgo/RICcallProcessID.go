// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICcallProcessID.h"
import "C"
import (
	"encoding/binary"
	"unsafe"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newRicCallProcessID(rcpID *e2ap_commondatatypes.RiccallProcessId) *C.RICcallProcessID_t {
	return newOctetString(rcpID.Value)
}

func decodeRicCallProcessIDBytes(rcpIDBytes []byte) *e2ap_commondatatypes.RiccallProcessId {
	rcpIDC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(rcpIDBytes[:8])))),
		size: C.ulong(binary.LittleEndian.Uint64(rcpIDBytes[8:])),
	}

	return decodeRicCallProcessID(&rcpIDC)
}

func decodeRicCallProcessID(rcpIDC *C.RICcallProcessID_t) *e2ap_commondatatypes.RiccallProcessId {
	result := e2ap_commondatatypes.RiccallProcessId{
		Value: decodeOctetString(rcpIDC),
	}

	return &result
}
