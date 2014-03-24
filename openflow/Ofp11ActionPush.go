package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionPush = 8

type Ofp11ActionPush struct {
	// msg_type uint16
	// len uint16
	// ethertype uint16
	// pad [2]byte
}

func (o Ofp11ActionPush) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionPush) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionPush) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionPush) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionPush) Ethertype(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp11ActionPush) SetEthertype(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
