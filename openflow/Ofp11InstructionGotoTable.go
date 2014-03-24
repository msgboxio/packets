package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11InstructionGotoTable = 8

type Ofp11InstructionGotoTable struct {
	// msg_type uint16
	// len uint16
	// table_id uint8
	// pad [3]byte
}

func (o Ofp11InstructionGotoTable) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11InstructionGotoTable) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11InstructionGotoTable) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11InstructionGotoTable) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11InstructionGotoTable) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o Ofp11InstructionGotoTable) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
