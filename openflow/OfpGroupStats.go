package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpGroupStats = 40

type OfpGroupStats struct {
	// length uint16
	// pad [2]byte
	// group_id uint32
	// ref_count uint32
	// pad2 [4]byte
	// packet_count uint64
	// byte_count uint64
	// duration_sec uint32
	// duration_nsec uint32
}

func (o OfpGroupStats) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpGroupStats) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpGroupStats) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpGroupStats) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpGroupStats) RefCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpGroupStats) SetRefCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpGroupStats) Pad2(b []byte) ([]byte, error) {
	return b[12:16], nil
}
func (o OfpGroupStats) SetPad2(b []byte, vn []byte) error {
	copy(b[12:16], vn)
	return nil
}
func (o OfpGroupStats) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpGroupStats) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o OfpGroupStats) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o OfpGroupStats) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
func (o OfpGroupStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o OfpGroupStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o OfpGroupStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o OfpGroupStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
