package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionNwTtl = 8

type OfpActionNwTtl struct {
	// msg_type uint16
	// len uint16
	// nw_ttl uint8
	// pad [3]byte
}

func (o OfpActionNwTtl) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionNwTtl) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionNwTtl) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionNwTtl) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionNwTtl) NwTtl(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o OfpActionNwTtl) SetNwTtl(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
