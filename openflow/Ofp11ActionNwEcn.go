package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionNwEcn = 8

type Ofp11ActionNwEcn struct {
	// msg_type uint16
	// len uint16
	// nw_ecn uint8
	// pad [3]byte
}

func (o Ofp11ActionNwEcn) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionNwEcn) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionNwEcn) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionNwEcn) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionNwEcn) NwEcn(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o Ofp11ActionNwEcn) SetNwEcn(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
