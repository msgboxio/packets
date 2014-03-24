package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11Instruction = 8

type Ofp11Instruction struct {
	// msg_type uint16
	// len uint16
	// pad [4]byte
}

func (o Ofp11Instruction) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11Instruction) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11Instruction) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11Instruction) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
