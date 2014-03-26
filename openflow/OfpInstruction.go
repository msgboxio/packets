package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpInstruction = 4

type OfpInstruction struct {
	// msg_type uint16
	// len uint16
}

func (o OfpInstruction) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpInstruction) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpInstruction) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpInstruction) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
