package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PacketIn = 10

type Ofp10PacketIn struct {
	// buffer_id uint32
	// total_len uint16
	// in_port uint16
	// reason uint8
	// pad uint8
	// data [0]byte
}

func (o Ofp10PacketIn) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp10PacketIn) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp10PacketIn) TotalLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp10PacketIn) SetTotalLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp10PacketIn) InPort(b []byte) (uint16, error) {
	return packets.ReadB16(b, 6)
}
func (o Ofp10PacketIn) SetInPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 6, vn)
}
func (o Ofp10PacketIn) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 8)
}
func (o Ofp10PacketIn) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 8, vn)
}
func (o Ofp10PacketIn) Data(b []byte) ([]byte, error) {
	return b[10:10], nil
}
func (o Ofp10PacketIn) SetData(b []byte, vn []byte) error {
	copy(b[10:10], vn)
	return nil
}
