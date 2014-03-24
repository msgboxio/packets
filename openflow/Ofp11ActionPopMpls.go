package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionPopMpls = 8

type Ofp11ActionPopMpls struct {
	// msg_type uint16
	// len uint16
	// ethertype uint16
	// pad [2]byte
}

func (o Ofp11ActionPopMpls) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionPopMpls) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionPopMpls) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionPopMpls) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionPopMpls) Ethertype(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp11ActionPopMpls) SetEthertype(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
