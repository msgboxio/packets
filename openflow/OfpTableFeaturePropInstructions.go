package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpTableFeaturePropInstructions = 4

type OfpTableFeaturePropInstructions struct {
	// msg_type uint16
	// length uint16
}

func (o OfpTableFeaturePropInstructions) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpTableFeaturePropInstructions) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpTableFeaturePropInstructions) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpTableFeaturePropInstructions) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
