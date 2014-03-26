package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpGroupDesc = 8

type OfpGroupDesc struct {
	// length uint16
	// msg_type uint8
	// pad uint8
	// group_id uint32
}

func (o OfpGroupDesc) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpGroupDesc) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpGroupDesc) MsgType(b []byte) (uint8, error) {
	return packets.ReadB8(b, 2)
}
func (o OfpGroupDesc) SetMsgType(b []byte, vn uint8) error {
	return packets.WriteB8(b, 2, vn)
}
func (o OfpGroupDesc) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpGroupDesc) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
