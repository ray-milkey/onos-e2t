// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func createRicServiceQueryMsg() (*e2ap_pdu_contents.RicserviceQuery, error) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	rsq, err := pdubuilder.CreateRicServiceQueryE2apPdu(1)
	if err != nil {
		return nil, err
	}
	rsq.GetInitiatingMessage().GetProcedureCode().GetRicServiceQuery().GetInitiatingMessage().SetRanFunctionsAccepted(rfAccepted)

	//if err := rsq.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating RicServiceQuery %s", err.Error())
	//}
	return rsq.GetInitiatingMessage().GetProcedureCode().GetRicServiceQuery().GetInitiatingMessage(), nil
}

func Test_xerEncodingRicServiceQuery(t *testing.T) {

	rsq, err := createRicServiceQueryMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdate PDU")

	xer, err := xerEncodeRicServiceQuery(rsq)
	assert.NilError(t, err)
	t.Logf("RicServiceQuery XER\n%s", string(xer))

	result, err := xerDecodeRicServiceQuery(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceQuery XER - decoded\n%v", result)
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}

func Test_perEncodingRicServiceQuery(t *testing.T) {

	rsq, err := createRicServiceQueryMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdate PDU")

	per, err := perEncodeRicServiceQuery(rsq)
	assert.NilError(t, err)
	t.Logf("RicServiceQuery PER\n%v", hex.Dump(per))

	result, err := perDecodeRicServiceQuery(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceQuery PER - decoded\n%v", result)
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[1].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsq.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}
