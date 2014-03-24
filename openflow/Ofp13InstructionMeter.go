package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13InstructionMeter = 8

type Ofp13InstructionMeter struct {
	// msg_type uint16
	// len uint16
	// meter_id uint32
}

func (o Ofp13InstructionMeter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13InstructionMeter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13InstructionMeter) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13InstructionMeter) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13InstructionMeter) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13InstructionMeter) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
