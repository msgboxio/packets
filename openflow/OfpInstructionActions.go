package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpInstructionActions = 8

type OfpInstructionActions struct {
	// msg_type uint16
	// len uint16
	// pad [4]byte
}

func (o OfpInstructionActions) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpInstructionActions) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpInstructionActions) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpInstructionActions) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
