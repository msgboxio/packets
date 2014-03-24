package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionGroup = 8

type Ofp11ActionGroup struct {
	// msg_type uint16
	// len uint16
	// group_id uint32
}

func (o Ofp11ActionGroup) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionGroup) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionGroup) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionGroup) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionGroup) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11ActionGroup) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
