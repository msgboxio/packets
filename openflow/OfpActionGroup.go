package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionGroup = 8

type OfpActionGroup struct {
	// msg_type uint16
	// len uint16
	// group_id uint32
}

func (o OfpActionGroup) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionGroup) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionGroup) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionGroup) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionGroup) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpActionGroup) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
