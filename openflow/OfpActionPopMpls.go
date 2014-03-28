package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpActionPopMpls = 8

type OfpActionPopMpls struct {
	// msg_type uint16
	// len uint16
	// ethertype uint16
	// pad [2]byte
}

func (o OfpActionPopMpls) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionPopMpls) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionPopMpls) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionPopMpls) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionPopMpls) Ethertype(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o OfpActionPopMpls) SetEthertype(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
