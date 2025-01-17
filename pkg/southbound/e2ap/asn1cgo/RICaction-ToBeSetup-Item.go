// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICaction-ToBeSetup-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeRicActionToBeSetupItem(ratbsi *e2appducontents.RicactionToBeSetupItem) ([]byte, error) {
	ratbsiCP, err := newRicActionToBeSetupItem(ratbsi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicActionToBeSetupItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICaction_ToBeSetup_Item, unsafe.Pointer(ratbsiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicActionToBeSetupItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicActionToBeSetupItem(ratbsi *e2appducontents.RicactionToBeSetupItem) ([]byte, error) {
	ratbsiCP, err := newRicActionToBeSetupItem(ratbsi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicActionToBeSetupItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICaction_ToBeSetup_Item, unsafe.Pointer(ratbsiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicActionToBeSetupItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicActionToBeSetupItem(bytes []byte) (*e2appducontents.RicactionToBeSetupItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICaction_ToBeSetup_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicActionToBeSetupItem((*C.RICaction_ToBeSetup_Item_t)(unsafePtr))
}

func perDecodeRicActionToBeSetupItem(bytes []byte) (*e2appducontents.RicactionToBeSetupItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICaction_ToBeSetup_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicActionToBeSetupItem((*C.RICaction_ToBeSetup_Item_t)(unsafePtr))
}

func newRicActionToBeSetupItem(ratbsItem *e2appducontents.RicactionToBeSetupItem) (*C.RICaction_ToBeSetup_Item_t, error) {
	ratC, err := newRicActionType(ratbsItem.GetRicActionType())
	if err != nil {
		return nil, fmt.Errorf("newRicActionType() %s", err.Error())
	}

	ratbsItemC := C.RICaction_ToBeSetup_Item_t{
		ricActionID:   *newRicActionID(ratbsItem.GetRicActionId()),
		ricActionType: *ratC,
		//ricActionDefinition: newRicActionDefinition(ratbsItem.GetRicActionDefinition()),
		//ricSubsequentAction: rsaC,
	}

	if ratbsItem.GetRicActionDefinition() != nil {
		ratbsItemC.ricActionDefinition = newRicActionDefinition(ratbsItem.GetRicActionDefinition())
	}
	if ratbsItem.RicSubsequentAction != nil {
		ratbsItemC.ricSubsequentAction, err = newRicSubsequentAction(ratbsItem.RicSubsequentAction)
		if err != nil {
			return nil, fmt.Errorf("newRicSubsequentAction() %s", err.Error())
		}
	}

	return &ratbsItemC, nil
}

func decodeRicActionToBeSetupItemBytes(bytes [56]byte) (*e2appducontents.RicactionToBeSetupItem, error) {

	rfiC := C.RICaction_ToBeSetup_Item_t{
		ricActionID:         C.long(binary.LittleEndian.Uint64(bytes[:8])),
		ricActionType:       C.long(binary.LittleEndian.Uint64(bytes[8:16])),
		ricActionDefinition: (*C.RICactionDefinition_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[16:24])))),
		ricSubsequentAction: (*C.struct_RICsubsequentAction)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[24:32])))),
	}

	return decodeRicActionToBeSetupItem(&rfiC)
}

func decodeRicActionToBeSetupItem(rfiC *C.RICaction_ToBeSetup_Item_t) (*e2appducontents.RicactionToBeSetupItem, error) {

	var err error
	rfi := e2appducontents.RicactionToBeSetupItem{
		RicActionId:   decodeRicActionID(&rfiC.ricActionID),
		RicActionType: decodeRicActionType(&rfiC.ricActionType),
		//RicActionDefinition: decodeRicActionDefinition(rfiC.ricActionDefinition),
		//RicSubsequentAction: rsa,
	}

	if rfiC.ricActionDefinition != nil {
		rfi.RicActionDefinition = decodeRicActionDefinition(rfiC.ricActionDefinition)
	}
	if rfiC.ricSubsequentAction != nil {
		rfi.RicSubsequentAction, err = decodeRicSubsequentAction(rfiC.ricSubsequentAction)
		if err != nil {
			return nil, fmt.Errorf("decodeRicSubsequentAction() %s", err.Error())
		}
	}

	return &rfi, nil
}
