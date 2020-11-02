// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalRIC-ID.h"
import "C"
import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// PerGlobalRicIDTOld - used for test only
// Deprecated: Do not use.
func PerGlobalRicIDTOld(e2srIE *e2ctypes.GlobalRIC_IDT) ([]byte, error) {
	rsrIEC, err := newGlobalRicIDOld(e2srIE)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalRIC_ID, unsafe.Pointer(rsrIEC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Deprecated: Do not use.
func newGlobalRicIDOld(id *e2ctypes.GlobalRIC_IDT) (*C.GlobalRIC_ID_t, error) {
	if len(id.PLMN_Identity) != 3 {
		return nil, fmt.Errorf("plmnID length is 3 - e2ap-v01.00.00.asn line 1105")
	}

	if id.Ric_ID.GetNumbits() != 20 {
		return nil, fmt.Errorf("ric-ID has to be 20 bits exactly - e2ap-v01.00.00.asn line 1076")
	}

	idC := C.GlobalRIC_ID_t{
		pLMN_Identity: *newOctetString(id.PLMN_Identity),
		ric_ID:        *newBitStringOld(id.Ric_ID),
	}

	return &idC, nil
}

func xerGlobalRicIDT(gr *e2apies.GlobalRicId) ([]byte, error) {
	grC, err := newGlobalRicID(gr)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalRIC_ID, unsafe.Pointer(grC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perGlobalRicIDT(gr *e2apies.GlobalRicId) ([]byte, error) {
	grC, err := newGlobalRicID(gr)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalRIC_ID, unsafe.Pointer(grC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newGlobalRicID(gr *e2apies.GlobalRicId) (*C.GlobalRIC_ID_t, error) {
	if len(gr.PLmnIdentity.Value) > 3 {
		return nil, fmt.Errorf("PLmnIdentity max length is 3 - e2ap-v01.00.00.asn line 1105")
	}
	if gr.RicId.Len != 20 {
		return nil, fmt.Errorf("ric-ID has to be 20 bits exactly - e2ap-v01.00.00.asn line 1076")
	}

	idC := C.GlobalRIC_ID_t{
		pLMN_Identity: *newOctetString(string(gr.PLmnIdentity.Value)),
		ric_ID:        *newBitString(gr.RicId),
	}

	return &idC, nil
}

func decodeGlobalRicID(grID *C.GlobalRIC_ID_t) (*e2apies.GlobalRicId, error) {
	result := new(e2apies.GlobalRicId)
	result.PLmnIdentity = new(e2ap_commondatatypes.PlmnIdentity)
	var err error
	result.PLmnIdentity.Value = []byte(decodeOctetString(&grID.pLMN_Identity))
	result.RicId, err = decodeBitString(&grID.ric_ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}