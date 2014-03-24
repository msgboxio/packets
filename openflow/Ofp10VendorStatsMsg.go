package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10VendorStatsMsg = 16

type Ofp10VendorStatsMsg struct {
	// osm ofp10_stats_msg
	// vendor uint32
}

func (o Ofp10VendorStatsMsg) Vendor(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp10VendorStatsMsg) SetVendor(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
