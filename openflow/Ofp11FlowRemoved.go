package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11FlowRemoved = 40

type Ofp11FlowRemoved struct {
	// cookie uint64
	// priority uint16
	// reason uint8
	// table_id uint8
	// duration_sec uint32
	// duration_nsec uint32
	// idle_timeout uint16
	// pad2 [2]byte
	// packet_count uint64
	// byte_count uint64
}

func (o Ofp11FlowRemoved) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o Ofp11FlowRemoved) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o Ofp11FlowRemoved) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp11FlowRemoved) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o Ofp11FlowRemoved) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 10)
}
func (o Ofp11FlowRemoved) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 10, vn)
}
func (o Ofp11FlowRemoved) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 11)
}
func (o Ofp11FlowRemoved) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 11, vn)
}
func (o Ofp11FlowRemoved) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp11FlowRemoved) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o Ofp11FlowRemoved) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o Ofp11FlowRemoved) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
func (o Ofp11FlowRemoved) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 20)
}
func (o Ofp11FlowRemoved) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 20, vn)
}
func (o Ofp11FlowRemoved) Pad2(b []byte) ([]byte, error) {
	return b[22:24], nil
}
func (o Ofp11FlowRemoved) SetPad2(b []byte, vn []byte) error {
	copy(b[22:24], vn)
	return nil
}
func (o Ofp11FlowRemoved) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 23)
}
func (o Ofp11FlowRemoved) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 23, vn)
}
func (o Ofp11FlowRemoved) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 31)
}
func (o Ofp11FlowRemoved) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 31, vn)
}
