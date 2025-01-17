// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunction-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeRanFunctionItem(rfi *e2appducontents.RanfunctionItem) ([]byte, error) {
	rfiCP := newRanFunctionItem(rfi)

	bytes, err := encodeXer(&C.asn_DEF_RANfunction_Item, unsafe.Pointer(rfiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanFunctionItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRanFunctionItem(rfi *e2appducontents.RanfunctionItem) ([]byte, error) {
	rfiCP := newRanFunctionItem(rfi)

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunction_Item, unsafe.Pointer(rfiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanFunctionItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRanFunctionItem(bytes []byte) (*e2appducontents.RanfunctionItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANfunction_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanFunctionItem((*C.RANfunction_Item_t)(unsafePtr))
}

func perDecodeRanFunctionItem(bytes []byte) (*e2appducontents.RanfunctionItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANfunction_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanFunctionItem((*C.RANfunction_Item_t)(unsafePtr))
}

func newRanFunctionItem(rfItem *e2appducontents.RanfunctionItem) *C.RANfunction_Item_t {
	rfItemC := C.RANfunction_Item_t{
		ranFunctionID:         newRanFunctionID(rfItem.GetRanFunctionId()),
		ranFunctionRevision:   *newRanFunctionRevision(rfItem.GetRanFunctionRevision()),
		ranFunctionDefinition: *newOctetString(rfItem.GetRanFunctionDefinition().GetValue()),
		ranFunctionOID:        *newRanFunctionOID(rfItem.GetRanFunctionOid()),
	}

	//if rfItem.GetRanFunctionOid() != nil {
	//	rfItemC.ranFunctionOID = newRanFunctionOID(rfItem.GetRanFunctionOid())
	//}

	return &rfItemC
}

func decodeRanFunctionItemBytes(bytes [120]byte) (*e2appducontents.RanfunctionItem, error) {
	size := binary.LittleEndian.Uint64(bytes[16:24])
	gobytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[8:16]))), C.int(size))
	sizeOID := binary.LittleEndian.Uint64(bytes[64:72])
	gobytesOID := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[56:64]))), C.int(sizeOID))

	rfiC := C.RANfunction_Item_t{
		ranFunctionID: C.long(binary.LittleEndian.Uint64(bytes[:8])),
		ranFunctionDefinition: C.OCTET_STRING_t{
			buf:  (*C.uchar)(C.CBytes(gobytes)),
			size: C.ulong(size),
		},
		ranFunctionRevision: C.long(binary.LittleEndian.Uint64(bytes[48:56])),
		//ranFunctionOID:      *(*C.PrintableString_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[56:64])))),
		ranFunctionOID: C.PrintableString_t{
			buf:  (*C.uchar)(C.CBytes(gobytesOID)),
			size: C.ulong(sizeOID),
		},
	}

	return decodeRanFunctionItem(&rfiC)
}

func decodeRanFunctionItem(rfiC *C.RANfunction_Item_t) (*e2appducontents.RanfunctionItem, error) {
	rfi := e2appducontents.RanfunctionItem{
		RanFunctionId: &e2apies.RanfunctionId{
			Value: decodeRanFunctionID(&rfiC.ranFunctionID).Value,
		},
		RanFunctionRevision: &e2apies.RanfunctionRevision{
			Value: decodeRanFunctionRevision(&rfiC.ranFunctionRevision).Value,
		},
		RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
			Value: decodeOctetString(&rfiC.ranFunctionDefinition),
		},
		RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
			Value: decodeRanFunctionOID(&rfiC.ranFunctionOID).Value,
		},
	}

	//if rfiC.ranFunctionOID != nil {
	//	rfi.RanFunctionOid = &e2ap_commondatatypes.RanfunctionOid{
	//		Value: decodeRanFunctionOID(rfiC.ranFunctionOID).Value,
	//	}
	//}
	return &rfi, nil
}
