package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionNwTos = 8

type OfpActionNwTos struct {
	// msg_type uint16
	// len uint16
	// nw_tos uint8
	// pad [3]byte
}

func (o OfpActionNwTos) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionNwTos) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionNwTos) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionNwTos) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionNwTos) NwTos(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o OfpActionNwTos) SetNwTos(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
