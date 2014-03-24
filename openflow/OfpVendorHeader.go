package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpVendorHeader = 12

type OfpVendorHeader struct {
	// header ofp_header
	// vendor uint32
}

func (o OfpVendorHeader) Vendor(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpVendorHeader) SetVendor(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
