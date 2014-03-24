package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionHeader = 8

type OfpActionHeader struct {
	// msg_type uint16
	// len uint16
	// pad [4]byte
}

func (o OfpActionHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionHeader) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionHeader) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
