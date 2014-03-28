package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpActionMplsTtl = 8

type OfpActionMplsTtl struct {
	// msg_type uint16
	// len uint16
	// mpls_ttl uint8
	// pad [3]byte
}

func (o OfpActionMplsTtl) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionMplsTtl) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionMplsTtl) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionMplsTtl) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionMplsTtl) MplsTtl(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o OfpActionMplsTtl) SetMplsTtl(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
