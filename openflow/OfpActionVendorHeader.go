package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionVendorHeader = 8

type OfpActionVendorHeader struct {
	// msg_type uint16
	// len uint16
	// vendor uint32
}

func (o OfpActionVendorHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionVendorHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionVendorHeader) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionVendorHeader) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionVendorHeader) Vendor(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpActionVendorHeader) SetVendor(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
