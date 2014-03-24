package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionNwAddr = 8

type OfpActionNwAddr struct {
	// msg_type uint16
	// len uint16
	// nw_addr uint32
}

func (o OfpActionNwAddr) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionNwAddr) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionNwAddr) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionNwAddr) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionNwAddr) NwAddr(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpActionNwAddr) SetNwAddr(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
