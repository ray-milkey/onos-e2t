// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-DU-ID.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeGnbDuID(gnbDuID *e2apies.GnbDuId) ([]byte, error) {
	gnbDuIDCP, err := newGnbDuID(gnbDuID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbDuID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GNB_DU_ID, unsafe.Pointer(gnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGnbDuID(gnbDuID *e2apies.GnbDuId) ([]byte, error) {
	gnbDuIDCP, err := newGnbDuID(gnbDuID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbDuID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_DU_ID, unsafe.Pointer(gnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGnbDuID(bytes []byte) (*e2apies.GnbDuId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGnbDuID((*C.GNB_DU_ID_t)(unsafePtr))
}

func perDecodeGnbDuID(bytes []byte) (*e2apies.GnbDuId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGnbDuID((*C.GNB_DU_ID_t)(unsafePtr))
}

func newGnbDuID(gnbDuID *e2apies.GnbDuId) (*C.GNB_DU_ID_t, error) {

	gnbDuIDC, err := newInteger(gnbDuID.Value)
	if err != nil {
		return nil, fmt.Errorf("newInteger() %s", err.Error())
	}

	return gnbDuIDC, nil
}

func decodeGnbDuID(gnbDuIDC *C.GNB_DU_ID_t) (*e2apies.GnbDuId, error) {

	gnbDuID := new(e2apies.GnbDuId)
	gnbDuIDValue, err := decodeInteger(gnbDuIDC)
	if err != nil {
		return nil, fmt.Errorf("decodeInteger() %s", err.Error())
	}
	gnbDuID.Value = gnbDuIDValue

	return gnbDuID, nil
}
