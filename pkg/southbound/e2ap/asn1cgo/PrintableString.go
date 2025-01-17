// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PrintableString.h"
import "C"

//import "unsafe"

func newPrintableString(msg string) *C.PrintableString_t {
	// PrintableString is defined via OctetString --> see PrintableString.h
	prntStrC := newOctetString([]byte(msg))

	return prntStrC
}

func decodePrintableString(octC *C.PrintableString_t) string {

	bytes := decodeOctetString(octC)
	return string(bytes)
}
