// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProcedureCode.h"
import "C"
import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newProcedureCode(pc *e2ap_commondatatypes.ProcedureCode) (C.ProcedureCode_t, error) {
	switch pcT := v2.ProcedureCodeT(pc.GetValue()); pcT {
	case v2.ProcedureCodeIDE2setup:
		return C.ProcedureCode_id_E2setup, nil
	case v2.ProcedureCodeIDErrorIndication:
		return C.ProcedureCode_id_ErrorIndication, nil
	case v2.ProcedureCodeIDReset:
		return C.ProcedureCode_id_Reset, nil
	case v2.ProcedureCodeIDRICcontrol:
		return C.ProcedureCode_id_RICcontrol, nil
	case v2.ProcedureCodeIDRICindication:
		return C.ProcedureCode_id_RICindication, nil
	case v2.ProcedureCodeIDRICserviceQuery:
		return C.ProcedureCode_id_RICserviceQuery, nil
	case v2.ProcedureCodeIDRICserviceUpdate:
		return C.ProcedureCode_id_RICserviceUpdate, nil
	case v2.ProcedureCodeIDRICsubscription:
		return C.ProcedureCode_id_RICsubscription, nil
	case v2.ProcedureCodeIDRICsubscriptionDelete:
		return C.ProcedureCode_id_RICsubscriptionDelete, nil
	case v2.ProcedureCodeIDRICsubscriptionDeleteRequired:
		return C.ProcedureCode_id_RICsubscriptionDeleteRequired, nil
	default:
		return 0, fmt.Errorf("unexpected procedure code %v", pcT)
	}
}

func decodeProcedureCode(pc C.ProcedureCode_t) *e2ap_commondatatypes.ProcedureCode {
	return &e2ap_commondatatypes.ProcedureCode{
		Value: int32(pc),
	}
}
