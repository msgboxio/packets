package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12ActionSetField = 8

type Ofp12ActionSetField struct {
	// msg_type uint16
	// len uint16
	// dst uint32
}

func (o Ofp12ActionSetField) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp12ActionSetField) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp12ActionSetField) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp12ActionSetField) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp12ActionSetField) Dst(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp12ActionSetField) SetDst(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
