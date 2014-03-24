package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11VendorStatsMsg = 20

type Ofp11VendorStatsMsg struct {
	// osm ofp11_stats_msg
	// vendor uint32
}

func (o Ofp11VendorStatsMsg) Vendor(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o Ofp11VendorStatsMsg) SetVendor(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
