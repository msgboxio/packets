package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpTableFeaturePropHeader = 4

type OfpTableFeaturePropHeader struct {
	// msg_type uint16
	// length uint16
}

func (o OfpTableFeaturePropHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpTableFeaturePropHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpTableFeaturePropHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpTableFeaturePropHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
