// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeConfigurationUpdateFailure.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure *e2ap_pdu_contents.E2NodeConfigurationUpdateFailure) ([]byte, error) {
	e2nodeConfigurationUpdateFailureCP, err := newE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeConfigurationUpdateFailure, unsafe.Pointer(e2nodeConfigurationUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure *e2ap_pdu_contents.E2NodeConfigurationUpdateFailure) ([]byte, error) {
	e2nodeConfigurationUpdateFailureCP, err := newE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeConfigurationUpdateFailure, unsafe.Pointer(e2nodeConfigurationUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeConfigurationUpdateFailure(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeConfigurationUpdateFailure((*C.E2nodeConfigurationUpdateFailure_t)(unsafePtr))
}

func perDecodeE2nodeConfigurationUpdateFailure(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeConfigurationUpdateFailure((*C.E2nodeConfigurationUpdateFailure_t)(unsafePtr))
}

func newE2nodeConfigurationUpdateFailure(e2cuf *e2ap_pdu_contents.E2NodeConfigurationUpdateFailure) (*C.E2nodeConfigurationUpdateFailure_t, error) {

	pIeC1751P19, err := newE2nodeConfigurationUpdateFailureIe(e2cuf.ProtocolIes)
	if err != nil {
		return nil, err
	}
	e2cufC := C.E2nodeConfigurationUpdateFailure_t{
		protocolIEs: *pIeC1751P19,
	}

	return &e2cufC, nil
}

func decodeE2nodeConfigurationUpdateFailure(e2cufC *C.E2nodeConfigurationUpdateFailure_t) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {

	pIEs, err := decodeE2nodeConfigurationUpdateFailureIes(&e2cufC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2cuf := e2ap_pdu_contents.E2NodeConfigurationUpdateFailure{
		ProtocolIes: pIEs,
	}

	return &e2cuf, nil
}
