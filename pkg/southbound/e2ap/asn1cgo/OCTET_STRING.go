// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "OCTET_STRING.h"
import "C"
import "unsafe"

// TODO: Change the argument to a []byte
func newOctetString(msg []byte) *C.OCTET_STRING_t {
	msgBytes := C.CBytes(msg)
	octStrC := C.OCTET_STRING_t{
		buf:  (*C.uchar)(msgBytes),
		size: C.ulong(len(msg)),
	}
	return &octStrC
}

func decodeOctetString(octC *C.OCTET_STRING_t) []byte {

	bytes := C.GoBytes(unsafe.Pointer(octC.buf), C.int(octC.size))
	return bytes
}
