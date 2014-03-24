package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11InstructionActions = 8

type Ofp11InstructionActions struct {
	// msg_type uint16
	// len uint16
	// pad [4]byte
}

func (o Ofp11InstructionActions) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11InstructionActions) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11InstructionActions) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11InstructionActions) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
