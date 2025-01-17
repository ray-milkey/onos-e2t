// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionSetupFailed-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2connectionSetupFailedItem(e2connectionSetupFailedItem *e2ap_pdu_contents.E2ConnectionSetupFailedItem) ([]byte, error) {
	e2connectionSetupFailedItemCP, err := newE2connectionSetupFailedItem(e2connectionSetupFailedItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionSetupFailedItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionSetupFailed_Item, unsafe.Pointer(e2connectionSetupFailedItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionSetupFailedItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionSetupFailedItem(e2connectionSetupFailedItem *e2ap_pdu_contents.E2ConnectionSetupFailedItem) ([]byte, error) {
	e2connectionSetupFailedItemCP, err := newE2connectionSetupFailedItem(e2connectionSetupFailedItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionSetupFailedItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionSetupFailed_Item, unsafe.Pointer(e2connectionSetupFailedItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionSetupFailedItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionSetupFailedItem(bytes []byte) (*e2ap_pdu_contents.E2ConnectionSetupFailedItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionSetupFailed_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionSetupFailedItem((*C.E2connectionSetupFailed_Item_t)(unsafePtr))
}

func perDecodeE2connectionSetupFailedItem(bytes []byte) (*e2ap_pdu_contents.E2ConnectionSetupFailedItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionSetupFailed_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionSetupFailedItem((*C.E2connectionSetupFailed_Item_t)(unsafePtr))
}

func newE2connectionSetupFailedItem(e2connectionSetupFailedItem *e2ap_pdu_contents.E2ConnectionSetupFailedItem) (*C.E2connectionSetupFailed_Item_t, error) {

	var err error
	e2connectionSetupFailedItemC := C.E2connectionSetupFailed_Item_t{}

	tnlInformationC, err := newTnlinformation(e2connectionSetupFailedItem.TnlInformation)
	if err != nil {
		return nil, fmt.Errorf("newTnlinformation() %s", err.Error())
	}

	causeC, err := newCause(e2connectionSetupFailedItem.Cause)
	if err != nil {
		return nil, fmt.Errorf("newCause() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2connectionSetupFailedItemC.tnlInformation = *tnlInformationC
	e2connectionSetupFailedItemC.cause = *causeC

	return &e2connectionSetupFailedItemC, nil
}

func decodeE2connectionSetupFailedItem(e2connectionSetupFailedItemC *C.E2connectionSetupFailed_Item_t) (*e2ap_pdu_contents.E2ConnectionSetupFailedItem, error) {

	var err error
	e2connectionSetupFailedItem := e2ap_pdu_contents.E2ConnectionSetupFailedItem{}

	e2connectionSetupFailedItem.TnlInformation, err = decodeTnlinformation(&e2connectionSetupFailedItemC.tnlInformation)
	if err != nil {
		return nil, fmt.Errorf("decodeTnlinformation() %s", err.Error())
	}

	e2connectionSetupFailedItem.Cause, err = decodeCause(&e2connectionSetupFailedItemC.cause)
	if err != nil {
		return nil, fmt.Errorf("decodeCause() %s", err.Error())
	}

	return &e2connectionSetupFailedItem, nil
}

func decodeE2connectionSetupFailedItemBytes(array [144]byte) (*e2ap_pdu_contents.E2ConnectionSetupFailedItem, error) {

	tnlAddrsize := binary.LittleEndian.Uint64(array[8:16])
	tnlAddrbitsUnused := int(binary.LittleEndian.Uint32(array[16:20]))
	tnlAddrbytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[:8]))), C.int(tnlAddrsize))

	tnlPortPtrC := (*C.BIT_STRING_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[48:56]))))

	e2csfItemC := C.E2connectionSetupFailed_Item_t{
		tnlInformation: C.TNLinformation_t{
			tnlAddress: C.BIT_STRING_t{
				buf:         (*C.uchar)(C.CBytes(tnlAddrbytes)),
				size:        C.ulong(tnlAddrsize),
				bits_unused: C.int(tnlAddrbitsUnused),
			},
			tnlPort: tnlPortPtrC,
		},
		cause: C.Cause_t{
			present: C.Cause_PR(binary.LittleEndian.Uint64(array[80:])),
		},
	}
	copy(e2csfItemC.cause.choice[:], array[88:96])

	return decodeE2connectionSetupFailedItem(&e2csfItemC)
}
