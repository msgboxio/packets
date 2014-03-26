package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpFlowStats = 56

type OfpFlowStats struct {
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
	// match ofp_match
}

func (o OfpFlowStats) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpFlowStats) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpFlowStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 2)
}
func (o OfpFlowStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 2, vn)
}
func (o OfpFlowStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpFlowStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpFlowStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpFlowStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpFlowStats) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 12)
}
func (o OfpFlowStats) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 12, vn)
}
func (o OfpFlowStats) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 14)
}
func (o OfpFlowStats) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 14, vn)
}
func (o OfpFlowStats) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 16)
}
func (o OfpFlowStats) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 16, vn)
}
func (o OfpFlowStats) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 18)
}
func (o OfpFlowStats) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 18, vn)
}
func (o OfpFlowStats) Pad2(b []byte) ([]byte, error) {
	return b[20:24], nil
}
func (o OfpFlowStats) SetPad2(b []byte, vn []byte) error {
	copy(b[20:24], vn)
	return nil
}
func (o OfpFlowStats) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o OfpFlowStats) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
func (o OfpFlowStats) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 32)
}
func (o OfpFlowStats) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 32, vn)
}
func (o OfpFlowStats) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o OfpFlowStats) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
