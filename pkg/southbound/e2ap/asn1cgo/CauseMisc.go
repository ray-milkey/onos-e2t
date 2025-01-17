// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseMisc.h"
import "C"
import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeCauseMisc(causeMisc *e2ap_ies.CauseMisc) ([]byte, error) {
	causeMiscCP, err := newCauseMisc(causeMisc)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseMisc, unsafe.Pointer(causeMiscCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseMisc() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseMisc(causeMisc *e2ap_ies.CauseMisc) ([]byte, error) {
	causeMiscCP, err := newCauseMisc(causeMisc)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseMisc, unsafe.Pointer(causeMiscCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseMisc() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseMisc(bytes []byte) (*e2ap_ies.CauseMisc, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseMisc)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseMisc((*C.CauseMisc_t)(unsafePtr))
}

func perDecodeCauseMisc(bytes []byte) (*e2ap_ies.CauseMisc, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseMisc)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseMisc((*C.CauseMisc_t)(unsafePtr))
}

func newCauseMisc(causeMisc *e2ap_ies.CauseMisc) (*C.CauseMisc_t, error) {
	var ret C.CauseMisc_t
	switch *causeMisc {
	case e2ap_ies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD:
		ret = C.CauseMisc_control_processing_overload
	case e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE:
		ret = C.CauseMisc_hardware_failure
	case e2ap_ies.CauseMisc_CAUSE_MISC_OM_INTERVENTION:
		ret = C.CauseMisc_om_intervention
	case e2ap_ies.CauseMisc_CAUSE_MISC_UNSPECIFIED:
		ret = C.CauseMisc_unspecified
	default:
		return nil, fmt.Errorf("unexpected CauseMisc %v", causeMisc)
	}

	return &ret, nil
}

func decodeCauseMisc(causeMiscC *C.CauseMisc_t) (*e2ap_ies.CauseMisc, error) {

	causeMisc := e2ap_ies.CauseMisc(int32(*causeMiscC))

	return &causeMisc, nil
}
