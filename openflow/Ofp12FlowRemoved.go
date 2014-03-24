package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12FlowRemoved = 40

type Ofp12FlowRemoved struct {
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
}

func (o Ofp12FlowRemoved) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o Ofp12FlowRemoved) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o Ofp12FlowRemoved) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp12FlowRemoved) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o Ofp12FlowRemoved) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 10)
}
func (o Ofp12FlowRemoved) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 10, vn)
}
func (o Ofp12FlowRemoved) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 11)
}
func (o Ofp12FlowRemoved) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 11, vn)
}
func (o Ofp12FlowRemoved) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp12FlowRemoved) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o Ofp12FlowRemoved) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o Ofp12FlowRemoved) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
func (o Ofp12FlowRemoved) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 20)
}
func (o Ofp12FlowRemoved) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 20, vn)
}
func (o Ofp12FlowRemoved) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 22)
}
func (o Ofp12FlowRemoved) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 22, vn)
}
func (o Ofp12FlowRemoved) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o Ofp12FlowRemoved) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
func (o Ofp12FlowRemoved) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 32)
}
func (o Ofp12FlowRemoved) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 32, vn)
}
