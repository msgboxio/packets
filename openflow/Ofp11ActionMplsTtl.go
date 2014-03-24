package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionMplsTtl = 8

type Ofp11ActionMplsTtl struct {
	// msg_type uint16
	// len uint16
	// mpls_ttl uint8
	// pad [3]byte
}

func (o Ofp11ActionMplsTtl) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionMplsTtl) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionMplsTtl) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionMplsTtl) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionMplsTtl) MplsTtl(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o Ofp11ActionMplsTtl) SetMplsTtl(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
