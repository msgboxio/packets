package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10FlowRemoved = 80

type Ofp10FlowRemoved struct {
	// match ofp10_match
	// cookie uint64
	// priority uint16
	// reason uint8
	// pad [1]byte
	// duration_sec uint32
	// duration_nsec uint32
	// idle_timeout uint16
	// pad2 [2]byte
	// packet_count uint64
	// byte_count uint64
}

func (o Ofp10FlowRemoved) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o Ofp10FlowRemoved) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
func (o Ofp10FlowRemoved) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 48)
}
func (o Ofp10FlowRemoved) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 48, vn)
}
func (o Ofp10FlowRemoved) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 50)
}
func (o Ofp10FlowRemoved) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 50, vn)
}
func (o Ofp10FlowRemoved) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 52)
}
func (o Ofp10FlowRemoved) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 52, vn)
}
func (o Ofp10FlowRemoved) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 56)
}
func (o Ofp10FlowRemoved) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 56, vn)
}
func (o Ofp10FlowRemoved) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 60)
}
func (o Ofp10FlowRemoved) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 60, vn)
}
func (o Ofp10FlowRemoved) Pad2(b []byte) ([]byte, error) {
	return b[62:64], nil
}
func (o Ofp10FlowRemoved) SetPad2(b []byte, vn []byte) error {
	copy(b[62:64], vn)
	return nil
}
func (o Ofp10FlowRemoved) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 63)
}
func (o Ofp10FlowRemoved) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 63, vn)
}
func (o Ofp10FlowRemoved) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 71)
}
func (o Ofp10FlowRemoved) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 71, vn)
}
