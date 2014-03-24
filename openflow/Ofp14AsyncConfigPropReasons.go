package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp14AsyncConfigPropReasons = 8

type Ofp14AsyncConfigPropReasons struct {
	// msg_type uint16
	// length uint16
	// mask uint32
}

func (o Ofp14AsyncConfigPropReasons) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp14AsyncConfigPropReasons) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp14AsyncConfigPropReasons) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp14AsyncConfigPropReasons) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp14AsyncConfigPropReasons) Mask(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp14AsyncConfigPropReasons) SetMask(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
