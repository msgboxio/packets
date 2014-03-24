package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11FlowStats = 48

type Ofp11FlowStats struct {
	// length uint16
	// table_id uint8
	// pad uint8
	// duration_sec uint32
	// duration_nsec uint32
	// priority uint16
	// idle_timeout uint16
	// hard_timeout uint16
	// flags uint16
	// pad2 [4]byte
	// cookie uint64
	// packet_count uint64
	// byte_count uint64
}

func (o Ofp11FlowStats) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11FlowStats) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11FlowStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 2)
}
func (o Ofp11FlowStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 2, vn)
}
func (o Ofp11FlowStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11FlowStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11FlowStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp11FlowStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp11FlowStats) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 12)
}
func (o Ofp11FlowStats) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 12, vn)
}
func (o Ofp11FlowStats) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 14)
}
func (o Ofp11FlowStats) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 14, vn)
}
func (o Ofp11FlowStats) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 16)
}
func (o Ofp11FlowStats) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 16, vn)
}
func (o Ofp11FlowStats) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 18)
}
func (o Ofp11FlowStats) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 18, vn)
}
func (o Ofp11FlowStats) Pad2(b []byte) ([]byte, error) {
	return b[20:24], nil
}
func (o Ofp11FlowStats) SetPad2(b []byte, vn []byte) error {
	copy(b[20:24], vn)
	return nil
}
func (o Ofp11FlowStats) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 21)
}
func (o Ofp11FlowStats) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 21, vn)
}
func (o Ofp11FlowStats) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 29)
}
func (o Ofp11FlowStats) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 29, vn)
}
func (o Ofp11FlowStats) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 37)
}
func (o Ofp11FlowStats) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 37, vn)
}
