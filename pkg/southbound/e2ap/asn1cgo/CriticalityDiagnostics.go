// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CriticalityDiagnostics.h"
//#include "RICrequestID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeCriticalityDiagnostics(cd *e2apies.CriticalityDiagnostics) ([]byte, error) {
	cdCP, err := newCriticalityDiagnostics(cd)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCriticalityDiagnostic() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_CriticalityDiagnostics, unsafe.Pointer(cdCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCriticalityDiagnostic() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCriticalityDiagnostics(cd *e2apies.CriticalityDiagnostics) ([]byte, error) {
	cdCP, err := newCriticalityDiagnostics(cd)
	if err != nil {
		return nil, fmt.Errorf("perEncodeCriticalityDiagnostic() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CriticalityDiagnostics, unsafe.Pointer(cdCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCriticalityDiagnostic() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCriticalityDiagnostics(bytes []byte) (*e2apies.CriticalityDiagnostics, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CriticalityDiagnostics)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCriticalityDiagnostics((*C.CriticalityDiagnostics_t)(unsafePtr))
}

func perDecodeCriticalityDiagnostics(bytes []byte) (*e2apies.CriticalityDiagnostics, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CriticalityDiagnostics)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCriticalityDiagnostics((*C.CriticalityDiagnostics_t)(unsafePtr))
}

func newCriticalityDiagnostics(cd *e2apies.CriticalityDiagnostics) (*C.CriticalityDiagnostics_t, error) {

	var err error
	cdC := C.CriticalityDiagnostics_t{}

	if cd.GetProcedureCode() != nil {
		pc, err := newProcedureCode(cd.GetProcedureCode())
		if err != nil {
			return nil, fmt.Errorf("newProcedureCode() %s", err.Error())
		}
		cdC.procedureCode = &pc
	}

	if cd.GetTriggeringMessage() != -1 {
		tm, err := newTriggeringMessage(cd.GetTriggeringMessage())
		if err != nil {
			return nil, fmt.Errorf("newTriggeringMessage() %s", err.Error())
		}
		cdC.triggeringMessage = &tm
	}

	if cd.GetProcedureCriticality() != -1 {
		c, err := criticalityToC(cd.GetProcedureCriticality())
		if err != nil {
			return nil, fmt.Errorf("criticalityToC() %s", err.Error())
		}
		cdC.procedureCriticality = &c
	}

	if cd.GetIEsCriticalityDiagnostics() != nil {
		cdC.iEsCriticalityDiagnostics, err = newCriticalityDiagnosticsIeList(cd.GetIEsCriticalityDiagnostics())
		if err != nil {
			return nil, fmt.Errorf("newCriticalityDiagnosticsIeList() %s", err.Error())
		}
	}

	if cd.GetRicRequestorId() != nil {
		cdC.ricRequestorID = newRicRequestID(cd.GetRicRequestorId())
	}

	return &cdC, nil
}

func decodeCriticalityDiagnosticsBytes(bytes []byte) (*e2apies.CriticalityDiagnostics, error) {

	//var ieCrDiag []byte
	//copy(bytes[32:], ieCrDiag)

	cdC := C.CriticalityDiagnostics_t{
		procedureCode:             (*C.long)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[:8])))),
		triggeringMessage:         (*C.long)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[8:])))),
		procedureCriticality:      (*C.long)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[16:])))),
		ricRequestorID:            (*C.RICrequestID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[24:])))),
		iEsCriticalityDiagnostics: (*C.CriticalityDiagnostics_IE_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[32:])))),
	}
	return decodeCriticalityDiagnostics(&cdC)
}

func decodeCriticalityDiagnostics(cdC *C.CriticalityDiagnostics_t) (*e2apies.CriticalityDiagnostics, error) {

	var err error
	ret := e2apies.CriticalityDiagnostics{}

	if cdC.iEsCriticalityDiagnostics != nil {
		ret.IEsCriticalityDiagnostics, err = decodeCriticalityDiagnosticsIeList(cdC.iEsCriticalityDiagnostics)
		if err != nil {
			return nil, fmt.Errorf("decodeCriticalityDiagnostics() %s", err.Error())
		}
	}

	if cdC.procedureCode != nil {
		ret.ProcedureCode = decodeProcedureCode(*cdC.procedureCode)
	}

	if cdC.triggeringMessage != nil {
		trMsg := decodeTriggeringMessage(*cdC.triggeringMessage)
		ret.TriggeringMessage = &trMsg
	}

	if cdC.procedureCriticality != nil {
		c := decodeCriticality(*cdC.procedureCriticality)
		ret.ProcedureCriticality = &c
	}

	if cdC.ricRequestorID != nil {
		ret.RicRequestorId = decodeRicRequestID(cdC.ricRequestorID)
	}

	return &ret, nil
}
