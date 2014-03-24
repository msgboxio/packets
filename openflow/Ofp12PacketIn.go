package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12PacketIn = 8

type Ofp12PacketIn struct {
	// buffer_id uint32
	// total_len uint16
	// reason uint8
	// table_id uint8
}

func (o Ofp12PacketIn) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp12PacketIn) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp12PacketIn) TotalLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp12PacketIn) SetTotalLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp12PacketIn) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 6)
}
func (o Ofp12PacketIn) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 6, vn)
}
func (o Ofp12PacketIn) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 7)
}
func (o Ofp12PacketIn) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 7, vn)
}
