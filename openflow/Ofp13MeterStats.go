package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterStats = 40

type Ofp13MeterStats struct {
	// meter_id uint32
	// len uint16
	// pad [6]byte
	// flow_count uint32
	// packet_in_count uint64
	// byte_in_count uint64
	// duration_sec uint32
	// duration_nsec uint32
}

func (o Ofp13MeterStats) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp13MeterStats) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp13MeterStats) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp13MeterStats) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp13MeterStats) FlowCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp13MeterStats) SetFlowCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o Ofp13MeterStats) PacketInCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o Ofp13MeterStats) SetPacketInCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o Ofp13MeterStats) ByteInCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o Ofp13MeterStats) SetByteInCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
func (o Ofp13MeterStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o Ofp13MeterStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o Ofp13MeterStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o Ofp13MeterStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
