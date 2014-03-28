package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpPacketIn = 32

type OfpPacketIn struct {
	// header ofp_header
	// buffer_id uint32
	// total_len uint16
	// reason uint8
	// table_id uint8
	// cookie uint64
	// match ofp_match
}

func (o OfpPacketIn) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpPacketIn) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpPacketIn) TotalLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 12)
}
func (o OfpPacketIn) SetTotalLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 12, vn)
}
func (o OfpPacketIn) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 14)
}
func (o OfpPacketIn) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 14, vn)
}
func (o OfpPacketIn) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 15)
}
func (o OfpPacketIn) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 15, vn)
}
func (o OfpPacketIn) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpPacketIn) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
