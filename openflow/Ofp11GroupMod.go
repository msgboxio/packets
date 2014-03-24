package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11GroupMod = 8

type Ofp11GroupMod struct {
	// command uint16
	// msg_type uint8
	// pad uint8
	// group_id uint32
}

func (o Ofp11GroupMod) Command(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11GroupMod) SetCommand(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11GroupMod) MsgType(b []byte) (uint8, error) {
	return packets.ReadB8(b, 2)
}
func (o Ofp11GroupMod) SetMsgType(b []byte, vn uint8) error {
	return packets.WriteB8(b, 2, vn)
}
func (o Ofp11GroupMod) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11GroupMod) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
