package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10FlowStats = 88

type Ofp10FlowStats struct {
	// length uint16
	// table_id uint8
	// pad uint8
	// match ofp10_match
	// duration_sec uint32
	// duration_nsec uint32
	// priority uint16
	// idle_timeout uint16
	// hard_timeout uint16
	// pad2 [6]byte
	// cookie uint64
	// packet_count uint64
	// byte_count uint64
}

func (o Ofp10FlowStats) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10FlowStats) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10FlowStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 2)
}
func (o Ofp10FlowStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 2, vn)
}
func (o Ofp10FlowStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 44)
}
func (o Ofp10FlowStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 44, vn)
}
func (o Ofp10FlowStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 48)
}
func (o Ofp10FlowStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 48, vn)
}
func (o Ofp10FlowStats) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 52)
}
func (o Ofp10FlowStats) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 52, vn)
}
func (o Ofp10FlowStats) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 54)
}
func (o Ofp10FlowStats) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 54, vn)
}
func (o Ofp10FlowStats) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 56)
}
func (o Ofp10FlowStats) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 56, vn)
}
func (o Ofp10FlowStats) Pad2(b []byte) ([]byte, error) {
	return b[58:64], nil
}
func (o Ofp10FlowStats) SetPad2(b []byte, vn []byte) error {
	copy(b[58:64], vn)
	return nil
}
func (o Ofp10FlowStats) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 59)
}
func (o Ofp10FlowStats) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 59, vn)
}
func (o Ofp10FlowStats) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 67)
}
func (o Ofp10FlowStats) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 67, vn)
}
func (o Ofp10FlowStats) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 75)
}
func (o Ofp10FlowStats) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 75, vn)
}
