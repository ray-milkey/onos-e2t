// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunctionsID-List.h"
//#include "ProtocolIE-SingleContainer.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeRanFunctionsIDList(rfl *e2appducontents.RanfunctionsIdList) ([]byte, error) {
	rflC, err := newRanFunctionsIDList(rfl)

	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RANfunctionsID_List, unsafe.Pointer(rflC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perEncodeRanFunctionsIDList(rfl *e2appducontents.RanfunctionsIdList) ([]byte, error) {
	rflC, err := newRanFunctionsIDList(rfl)

	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunctionsID_List, unsafe.Pointer(rflC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func xerDecodeRanFunctionsIDList(xer []byte) (*e2appducontents.RanfunctionsIdList, error) {
	unsafePtr, err := decodeXer(xer, &C.asn_DEF_RANfunctionsID_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	rflC := (*C.RANfunctionsID_List_t)(unsafePtr)
	return decodeRanFunctionsIDList(rflC)
}

func perDecodeRanFunctionsIDList(per []byte) (*e2appducontents.RanfunctionsIdList, error) {
	unsafePtr, err := decodePer(per, len(per), &C.asn_DEF_RANfunctionsID_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	rflC := (*C.RANfunctionsID_List_t)(unsafePtr)
	return decodeRanFunctionsIDList(rflC)
}

func newRanFunctionsIDList(rfIDl *e2appducontents.RanfunctionsIdList) (*C.RANfunctionsID_List_t, error) {
	rfIDlC := new(C.RANfunctionsID_List_t)

	for _, rfID := range rfIDl.GetValue() {
		rfIDC, err := newRanFunctionIDItemIesSingleContainer(rfID)
		if err != nil {
			return nil, fmt.Errorf("error on newRanFunctionIDItemIesSingleContainer() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(rfIDlC), unsafe.Pointer(rfIDC)); err != nil {
			return nil, err
		}
	}

	return rfIDlC, nil
}

func decodeRanFunctionsIDListBytes(ranFunctionIDListChoice [112]byte) (*e2appducontents.RanfunctionsIdList, error) {
	array := (**C.struct_ProtocolIE_SingleContainer)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(ranFunctionIDListChoice[0:8]))))
	count := C.int(binary.LittleEndian.Uint32(ranFunctionIDListChoice[8:12]))
	size := C.int(binary.LittleEndian.Uint32(ranFunctionIDListChoice[12:16]))

	rfIDlC := C.RANfunctionsID_List_t{
		list: C.struct___150{
			array: array,
			size:  size,
			count: count,
		},
	}

	return decodeRanFunctionsIDList(&rfIDlC)
}

func decodeRanFunctionsIDList(rfIDlC *C.RANfunctionsID_List_t) (*e2appducontents.RanfunctionsIdList, error) {
	rfIDl := e2appducontents.RanfunctionsIdList{
		Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
	}

	ieCount := int(rfIDlC.list.count)
	//fmt.Printf("RanFunctionIDListC %T List %T %v Array %T %v Deref %v\n", rflC, rflC.list, rflC.list, rflC.list.array, *rflC.list.array, *(rflC.list.array))
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*rfIDlC.list.array)) * uintptr(i)
		rfIDiIeC := *(**C.ProtocolIE_SingleContainer_1911P15_t)(unsafe.Pointer(uintptr(unsafe.Pointer(rfIDlC.list.array)) + offset))
		//fmt.Printf("Value %T %p %v\n", rfIDiIeC, rfIDiIeC, rfIDiIeC)
		rfIDiIe, err := decodeRanFunctionIDItemIesSingleContainer(rfIDiIeC)
		if err != nil {
			return nil, fmt.Errorf("decodeRanFunctionIDItemIesSingleContainer() %s", err.Error())
		}
		rfIDl.Value = append(rfIDl.Value, rfIDiIe)
	}

	return &rfIDl, nil
}
