package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionPush = 8

type OfpActionPush struct {
	// msg_type uint16
	// len uint16
	// ethertype uint16
	// pad [2]byte
}

func (o OfpActionPush) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionPush) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionPush) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionPush) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionPush) Ethertype(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o OfpActionPush) SetEthertype(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
