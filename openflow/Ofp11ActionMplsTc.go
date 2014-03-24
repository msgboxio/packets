package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionMplsTc = 8

type Ofp11ActionMplsTc struct {
	// msg_type uint16
	// len uint16
	// mpls_tc uint8
	// pad [3]byte
}

func (o Ofp11ActionMplsTc) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionMplsTc) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionMplsTc) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionMplsTc) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionMplsTc) MplsTc(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o Ofp11ActionMplsTc) SetMplsTc(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
