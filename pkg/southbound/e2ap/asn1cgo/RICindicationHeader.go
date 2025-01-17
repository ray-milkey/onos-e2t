// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICindicationHeader.h"
import "C"
import (
	"encoding/binary"
	"unsafe"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newRicIndicationHeader(rih *e2ap_commondatatypes.RicindicationHeader) *C.RICindicationHeader_t {
	return newOctetString(rih.GetValue())
}

func decodeRicIndicationHeaderBytes(rihBytes []byte) *e2ap_commondatatypes.RicindicationHeader {
	rihC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(rihBytes[:8])))),
		size: C.ulong(binary.LittleEndian.Uint64(rihBytes[8:])),
	}
	return decodeRicIndicationHeader(&rihC)
}

func decodeRicIndicationHeader(rihC *C.RICindicationHeader_t) *e2ap_commondatatypes.RicindicationHeader {
	result := e2ap_commondatatypes.RicindicationHeader{
		Value: decodeOctetString(rihC),
	}

	return &result
}
