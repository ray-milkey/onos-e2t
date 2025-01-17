// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseRICservice.h"
import "C"
import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeCauseRicservice(causeRicservice *e2ap_ies.CauseRicservice) ([]byte, error) {
	causeRicserviceCP, err := newCauseRicservice(causeRicservice)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseRICservice, unsafe.Pointer(causeRicserviceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseRicservice() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseRicservice(causeRicservice *e2ap_ies.CauseRicservice) ([]byte, error) {
	causeRicserviceCP, err := newCauseRicservice(causeRicservice)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseRICservice, unsafe.Pointer(causeRicserviceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseRicservice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseRicservice(bytes []byte) (*e2ap_ies.CauseRicservice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseRICservice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseRicservice((*C.CauseRICservice_t)(unsafePtr))
}

func perDecodeCauseRicservice(bytes []byte) (*e2ap_ies.CauseRicservice, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseRICservice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseRicservice((*C.CauseRICservice_t)(unsafePtr))
}

func newCauseRicservice(causeRicservice *e2ap_ies.CauseRicservice) (*C.CauseRICservice_t, error) {
	var ret C.CauseRICservice_t
	switch *causeRicservice {
	case e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RAN_FUNCTION_NOT_SUPPORTED:
		ret = C.CauseRICservice_ran_function_not_supported
	case e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS:
		ret = C.CauseRICservice_excessive_functions
	case e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT:
		ret = C.CauseRICservice_ric_resource_limit
	default:
		return nil, fmt.Errorf("unexpected CauseRicservice %v", causeRicservice)
	}

	return &ret, nil
}

func decodeCauseRicservice(causeRicserviceC *C.CauseRICservice_t) (*e2ap_ies.CauseRicservice, error) {

	causeRicservice := e2ap_ies.CauseRicservice(int32(*causeRicserviceC))

	return &causeRicservice, nil
}
