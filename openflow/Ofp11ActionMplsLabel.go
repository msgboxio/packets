package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionMplsLabel = 8

type Ofp11ActionMplsLabel struct {
	// msg_type uint16
	// len uint16
	// mpls_label uint32
}

func (o Ofp11ActionMplsLabel) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionMplsLabel) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionMplsLabel) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionMplsLabel) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionMplsLabel) MplsLabel(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11ActionMplsLabel) SetMplsLabel(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
