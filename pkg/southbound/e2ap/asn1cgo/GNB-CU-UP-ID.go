// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-CU-UP-ID.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeGnbCuUpID(gnbCuUpID *e2apies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIDCP, err := newGnbCuUpID(gnbCuUpID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbCuUpID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbCuUpID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGnbCuUpID(gnbCuUpID *e2apies.GnbCuUpId) ([]byte, error) {
	gnbCuUpIDCP, err := newGnbCuUpID(gnbCuUpID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbCuUpID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_CU_UP_ID, unsafe.Pointer(gnbCuUpIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbCuUpID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGnbCuUpID(bytes []byte) (*e2apies.GnbCuUpId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGnbCuUpID((*C.GNB_CU_UP_ID_t)(unsafePtr))
}

func perDecodeGnbCuUpID(bytes []byte) (*e2apies.GnbCuUpId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGnbCuUpID((*C.GNB_CU_UP_ID_t)(unsafePtr))
}

func newGnbCuUpID(gnbCuUpID *e2apies.GnbCuUpId) (*C.GNB_CU_UP_ID_t, error) {

	gnbCuUpIDC, err := newInteger(gnbCuUpID.Value)
	if err != nil {
		return nil, fmt.Errorf("newInteger() %s", err.Error())
	}

	return gnbCuUpIDC, nil
}

func decodeGnbCuUpID(gnbCuUpIDC *C.GNB_CU_UP_ID_t) (*e2apies.GnbCuUpId, error) {

	gnbCuUpID := new(e2apies.GnbCuUpId)
	gnbCuUpIDValue, err := decodeInteger(gnbCuUpIDC)
	if err != nil {
		return nil, fmt.Errorf("decodeInteger() %s", err.Error())
	}
	gnbCuUpID.Value = gnbCuUpIDValue

	return gnbCuUpID, nil
}
