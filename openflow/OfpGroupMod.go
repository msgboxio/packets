package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpGroupMod = 16

type OfpGroupMod struct {
	// header ofp_header
	// command uint16
	// msg_type uint8
	// pad uint8
	// group_id uint32
}

func (o OfpGroupMod) Command(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpGroupMod) SetCommand(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpGroupMod) MsgType(b []byte) (uint8, error) {
	return packets.ReadB8(b, 10)
}
func (o OfpGroupMod) SetMsgType(b []byte, vn uint8) error {
	return packets.WriteB8(b, 10, vn)
}
func (o OfpGroupMod) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpGroupMod) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
