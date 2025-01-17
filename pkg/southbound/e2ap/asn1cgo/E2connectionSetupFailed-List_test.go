// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createE2connectionSetupFailedListMsg() (*e2ap_pdu_contents.E2ConnectionSetupFailedList, error) {

	e2connectionSetupFailedList := e2ap_pdu_contents.E2ConnectionSetupFailedList{
		Value: make([]*e2ap_pdu_contents.E2ConnectionSetupFailedItemIes, 0),
	}

	bs1 := &asn1.BitString{
		Value: []byte{0xab, 0xbc, 0xcd, 0xde, 0xef, 0xf5, 0xd6, 0xb7},
		Len:   64,
	}

	bs2 := &asn1.BitString{
		Value: []byte{0xcd, 0x9b},
		Len:   16,
	}

	tnlinformation := e2ap_ies.Tnlinformation{
		TnlAddress: bs1,
		TnlPort:    bs2,
	}

	item := &e2ap_pdu_contents.E2ConnectionSetupFailedItemIes{
		Id:          int32(v2.ProtocolIeIDE2connectionSetupFailedItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_pdu_contents.E2ConnectionSetupFailedItem{
			TnlInformation: &tnlinformation,
			Cause: &e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_RicService{
					RicService: e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
				},
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2connectionSetupFailedList.Value = append(e2connectionSetupFailedList.Value, item)

	//if err := e2connectionSetupFailedList.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2connectionSetupFailedList %s", err.Error())
	//}
	return &e2connectionSetupFailedList, nil
}

func Test_xerEncodingE2connectionSetupFailedList(t *testing.T) {

	e2connectionSetupFailedList, err := createE2connectionSetupFailedListMsg()
	assert.NilError(t, err, "Error creating E2connectionSetupFailedList PDU")

	xer, err := xerEncodeE2connectionSetupFailedList(e2connectionSetupFailedList)
	assert.NilError(t, err)
	assert.Equal(t, 756, len(xer))
	t.Logf("E2connectionSetupFailedList XER\n%s", string(xer))

	result, err := xerDecodeE2connectionSetupFailedList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionSetupFailedList XER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetCause().GetRicService(), result.GetValue()[0].GetValue().GetCause().GetRicService())
	assert.Equal(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
}

func Test_perEncodingE2connectionSetupFailedList(t *testing.T) {

	e2connectionSetupFailedList, err := createE2connectionSetupFailedListMsg()
	assert.NilError(t, err, "Error creating E2connectionSetupFailedList PDU")

	per, err := perEncodeE2connectionSetupFailedList(e2connectionSetupFailedList)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(per))
	t.Logf("E2connectionSetupFailedList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionSetupFailedList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionSetupFailedList PER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetCause().GetRicService(), result.GetValue()[0].GetValue().GetCause().GetRicService())
	assert.Equal(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionSetupFailedList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
}
