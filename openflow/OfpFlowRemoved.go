package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpFlowRemoved = 56

type OfpFlowRemoved struct {
	// header ofp_header
	// cookie uint64
	// priority uint16
	// reason uint8
	// table_id uint8
	// duration_sec uint32
	// duration_nsec uint32
	// idle_timeout uint16
	// hard_timeout uint16
	// packet_count uint64
	// byte_count uint64
	// match ofp_match
}

func (o OfpFlowRemoved) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpFlowRemoved) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpFlowRemoved) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 16)
}
func (o OfpFlowRemoved) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 16, vn)
}
func (o OfpFlowRemoved) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 18)
}
func (o OfpFlowRemoved) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 18, vn)
}
func (o OfpFlowRemoved) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 19)
}
func (o OfpFlowRemoved) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 19, vn)
}
func (o OfpFlowRemoved) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 20)
}
func (o OfpFlowRemoved) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 20, vn)
}
func (o OfpFlowRemoved) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o OfpFlowRemoved) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
func (o OfpFlowRemoved) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 28)
}
func (o OfpFlowRemoved) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 28, vn)
}
func (o OfpFlowRemoved) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 30)
}
func (o OfpFlowRemoved) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 30, vn)
}
func (o OfpFlowRemoved) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 32)
}
func (o OfpFlowRemoved) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 32, vn)
}
func (o OfpFlowRemoved) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o OfpFlowRemoved) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
