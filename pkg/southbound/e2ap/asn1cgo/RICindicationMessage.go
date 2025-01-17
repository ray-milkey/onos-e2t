// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICindicationMessage.h"
import "C"
import (
	"encoding/binary"
	"unsafe"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newRicIndicationMessage(rih *e2ap_commondatatypes.RicindicationMessage) *C.RICindicationMessage_t {
	return newOctetString(rih.GetValue())
}

func decodeRicIndicationMessageBytes(rimBytes []byte) *e2ap_commondatatypes.RicindicationMessage {
	rimC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(rimBytes[:8])))),
		size: C.ulong(binary.LittleEndian.Uint64(rimBytes[8:])),
	}
	return decodeRicIndicationMessage(&rimC)
}

func decodeRicIndicationMessage(rihC *C.RICindicationMessage_t) *e2ap_commondatatypes.RicindicationMessage {
	result := e2ap_commondatatypes.RicindicationMessage{
		Value: decodeOctetString(rihC),
	}

	return &result
}
