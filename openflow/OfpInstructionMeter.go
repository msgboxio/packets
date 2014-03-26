package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpInstructionMeter = 8

type OfpInstructionMeter struct {
	// msg_type uint16
	// len uint16
	// meter_id uint32
}

func (o OfpInstructionMeter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpInstructionMeter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpInstructionMeter) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpInstructionMeter) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpInstructionMeter) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpInstructionMeter) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
