package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpInstructionGotoTable = 8

type OfpInstructionGotoTable struct {
	// msg_type uint16
	// len uint16
	// table_id uint8
	// pad [3]byte
}

func (o OfpInstructionGotoTable) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpInstructionGotoTable) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpInstructionGotoTable) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpInstructionGotoTable) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpInstructionGotoTable) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o OfpInstructionGotoTable) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
