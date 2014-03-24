package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11PacketIn = 16

type Ofp11PacketIn struct {
	// buffer_id uint32
	// in_port uint32
	// in_phy_port uint32
	// total_len uint16
	// reason uint8
	// table_id uint8
}

func (o Ofp11PacketIn) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11PacketIn) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11PacketIn) InPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11PacketIn) SetInPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11PacketIn) InPhyPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp11PacketIn) SetInPhyPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp11PacketIn) TotalLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 12)
}
func (o Ofp11PacketIn) SetTotalLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 12, vn)
}
func (o Ofp11PacketIn) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 14)
}
func (o Ofp11PacketIn) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 14, vn)
}
func (o Ofp11PacketIn) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 15)
}
func (o Ofp11PacketIn) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 15, vn)
}
