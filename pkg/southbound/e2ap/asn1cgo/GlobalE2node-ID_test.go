// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createGlobalE2nodeIDGNb() *e2apies.GlobalE2NodeId {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_GNb{
			GNb: &e2apies.GlobalE2NodeGnbId{
				GlobalGNbId: &e2apies.GlobalgNbId{
					PlmnId: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{0x01, 0x02, 0x03},
					},
					GnbId: &e2apies.GnbIdChoice{
						GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
							GnbId: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x9c},
								Len:   22,
							},
						},
					},
				},
				GNbCuUpId: &e2apies.GnbCuUpId{
					Value: 21,
				},
				GNbDuId: &e2apies.GnbDuId{
					Value: 13,
				},
			},
		},
	}
}

func createGlobalE2nodeIDenGNb() *e2apies.GlobalE2NodeId {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_EnGNb{
			EnGNb: &e2apies.GlobalE2NodeEnGnbId{
				GlobalEnGNbId: &e2apies.GlobalenGnbId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{0x01, 0x02, 0x03},
					},
					GNbId: &e2apies.EngnbId{
						EngnbId: &e2apies.EngnbId_GNbId{
							GNbId: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x9c},
								Len:   22,
							},
						},
					},
				},
			},
		},
	}
}

func createGlobalE2nodeIDENb() *e2apies.GlobalE2NodeId {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_ENb{
			ENb: &e2apies.GlobalE2NodeEnbId{
				GlobalENbId: &e2apies.GlobalEnbId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{0x01, 0x02, 0x03},
					},
					ENbId: &e2apies.EnbId{
						EnbId: &e2apies.EnbId_MacroENbId{
							MacroENbId: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x90},
								Len:   20,
							},
						},
					},
				},
			},
		},
	}
}

func createGlobalE2nodeIDngENb() *e2apies.GlobalE2NodeId {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_NgENb{
			NgENb: &e2apies.GlobalE2NodeNgEnbId{
				GlobalNgENbId: &e2apies.GlobalngeNbId{
					PlmnId: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{0x01, 0x02, 0x03},
					},
					EnbId: &e2apies.EnbIdChoice{
						EnbIdChoice: &e2apies.EnbIdChoice_EnbIdLongmacro{
							EnbIdLongmacro: &asn1.BitString{
								Value: []byte{0xd4, 0xbc, 0x98},
								Len:   21,
							},
						},
					},
				},
			},
		},
	}
}

func Test_xerDecodeGlobalE2nodeID(t *testing.T) {

	ge2n := createGlobalE2nodeIDGNb()

	xer, err := xerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (GNb) XER\n%s", xer)

	// Now reverse the XER
	ge2nReversed, err := xerDecodeGlobalE2nodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalE2nodeID (GNb) decoded from XER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), ge2nReversed.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.DeepEqual(t, ge2n.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), ge2nReversed.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())

	ge2n = createGlobalE2nodeIDenGNb()

	xer, err = xerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (en-GNb) XER\n%s", xer)

	// Now reverse the XER
	ge2nReversed, err = xerDecodeGlobalE2nodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalE2nodeID (en-GNb) decoded from XER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), ge2nReversed.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
	assert.DeepEqual(t, ge2n.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), ge2nReversed.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())

	ge2n = createGlobalE2nodeIDENb()

	xer, err = xerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (ENb) XER\n%s", xer)

	// Now reverse the XER
	ge2nReversed, err = xerDecodeGlobalE2nodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalE2nodeID (ENb) decoded from XER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetLen(), ge2nReversed.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetLen())
	assert.DeepEqual(t, ge2n.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetValue(), ge2nReversed.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetValue())

	ge2n = createGlobalE2nodeIDngENb()

	xer, err = xerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (ng-ENb) XER\n%s", xer)

	// Now reverse the XER
	ge2nReversed, err = xerDecodeGlobalE2nodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalE2nodeID (ng-ENb) decoded from XER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetLen(), ge2nReversed.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetLen())
	assert.DeepEqual(t, ge2n.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetValue(), ge2nReversed.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetValue())
}

func Test_perDecodeGlobalE2nodeID(t *testing.T) {

	ge2n := createGlobalE2nodeIDGNb()

	per, err := PerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (GNb) PER\n%s", hex.Dump(per))

	// Now reverse the PER
	ge2nReversedFromPer, err := PerDecodeGlobalE2nodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversedFromPer != nil)
	t.Logf("GlobalE2nodeID (GNb) decoded from PER is \n%v", ge2nReversedFromPer)
	assert.Equal(t, ge2n.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), ge2nReversedFromPer.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	assert.DeepEqual(t, ge2n.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), ge2nReversedFromPer.GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())

	ge2n = createGlobalE2nodeIDenGNb()

	per, err = PerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (en-GNb) PER\n%s", hex.Dump(per))

	// Now reverse the PER
	ge2nReversedFromPer, err = PerDecodeGlobalE2nodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversedFromPer != nil)
	t.Logf("GlobalE2nodeID (en-GNb) decoded from PER is \n%v", ge2nReversedFromPer)
	assert.Equal(t, ge2n.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), ge2nReversedFromPer.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
	assert.DeepEqual(t, ge2n.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), ge2nReversedFromPer.GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())

	ge2n = createGlobalE2nodeIDENb()

	per, err = PerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (ENb) PER\n%s", hex.Dump(per))

	// Now reverse the PER
	ge2nReversedFromPer, err = PerDecodeGlobalE2nodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversedFromPer != nil)
	t.Logf("GlobalE2nodeID (ENb) decoded from PER is \n%v", ge2nReversedFromPer)
	assert.Equal(t, ge2n.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetLen(), ge2nReversedFromPer.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetLen())
	assert.DeepEqual(t, ge2n.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetValue(), ge2nReversedFromPer.GetENb().GetGlobalENbId().GetENbId().GetMacroENbId().GetValue())

	ge2n = createGlobalE2nodeIDngENb()

	per, err = PerEncodeGlobalE2nodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodeID (ng-ENb) PER\n%s", hex.Dump(per))

	// Now reverse the PER
	ge2nReversedFromPer, err = PerDecodeGlobalE2nodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversedFromPer != nil)
	t.Logf("GlobalE2nodeID (ng-ENb) decoded from PER is \n%v", ge2nReversedFromPer)
	assert.Equal(t, ge2n.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetLen(), ge2nReversedFromPer.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetLen())
	assert.DeepEqual(t, ge2n.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetValue(), ge2nReversedFromPer.GetNgENb().GetGlobalNgENbId().GetEnbId().GetEnbIdLongmacro().GetValue())
}
