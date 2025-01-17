// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func Test_RicActionsToBeSetupList(t *testing.T) {
	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[100] = types.RicActionDef{
		RicActionID:         100,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS,
		RicActionDefinition: []byte{0x11, 0x22},
	}

	ricActionsToBeSetup[200] = types.RicActionDef{
		RicActionID:         200,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
		RicActionDefinition: []byte{0x33, 0x44},
	}

	rsr, err := pdubuilder.NewRicSubscriptionRequest(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

	ricSubDetails := rsr.GetProtocolIes().GetE2ApProtocolIes30().GetValue()

	xer, err := xerEncodeRicActionsToBeSetupList(ricSubDetails.GetRicActionToBeSetupList())
	assert.NilError(t, err)
	t.Logf("RicActionToBeSetupList XER\n%s", xer)

	per, err := perEncodeRicActionsToBeSetupList(ricSubDetails.GetRicActionToBeSetupList())
	assert.NilError(t, err)
	t.Logf("RicActionToBeSetupList PER\n%v", hex.Dump(per))

	// Now reverse it
	ratbsL, err := xerDecodeRicActionsToBeSetupList(xer)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(ratbsL.GetValue()))

	for _, raTbsItem := range ratbsL.GetValue() {
		assert.Equal(t, int32(v2.ProtocolIeIDRicactionToBeSetupItem), raTbsItem.GetId())
		assert.Equal(t, e2apies.RicactionType_RICACTION_TYPE_INSERT, raTbsItem.GetValue().GetRicActionType())
		assert.Equal(t, e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE, raTbsItem.GetValue().GetRicSubsequentAction().GetRicSubsequentActionType())
		switch raID := raTbsItem.GetValue().GetRicActionId().GetValue(); raID {
		case 100:
			assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS,
				raTbsItem.GetValue().GetRicSubsequentAction().GetRicTimeToWait())
		case 200:
			assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
				raTbsItem.GetValue().GetRicSubsequentAction().GetRicTimeToWait())
		default:
			assert.Assert(t, false, "Unexpected RicActionID %d", raID)
		}
	}
}
