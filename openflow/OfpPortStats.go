package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpPortStats = 112

type OfpPortStats struct {
	// port_no uint32
	// pad [4]byte
	// rx_packets uint64
	// tx_packets uint64
	// rx_bytes uint64
	// tx_bytes uint64
	// rx_dropped uint64
	// tx_dropped uint64
	// rx_errors uint64
	// tx_errors uint64
	// rx_frame_err uint64
	// rx_over_err uint64
	// rx_crc_err uint64
	// collisions uint64
	// duration_sec uint32
	// duration_nsec uint32
}

func (o OfpPortStats) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpPortStats) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpPortStats) RxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpPortStats) SetRxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpPortStats) TxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpPortStats) SetTxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o OfpPortStats) RxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o OfpPortStats) SetRxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
func (o OfpPortStats) TxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 32)
}
func (o OfpPortStats) SetTxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 32, vn)
}
func (o OfpPortStats) RxDropped(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o OfpPortStats) SetRxDropped(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
func (o OfpPortStats) TxDropped(b []byte) (uint64, error) {
	return packets.ReadB64(b, 48)
}
func (o OfpPortStats) SetTxDropped(b []byte, vn uint64) error {
	return packets.WriteB64(b, 48, vn)
}
func (o OfpPortStats) RxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 56)
}
func (o OfpPortStats) SetRxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 56, vn)
}
func (o OfpPortStats) TxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 64)
}
func (o OfpPortStats) SetTxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 64, vn)
}
func (o OfpPortStats) RxFrameErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 72)
}
func (o OfpPortStats) SetRxFrameErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 72, vn)
}
func (o OfpPortStats) RxOverErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 80)
}
func (o OfpPortStats) SetRxOverErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 80, vn)
}
func (o OfpPortStats) RxCrcErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 88)
}
func (o OfpPortStats) SetRxCrcErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 88, vn)
}
func (o OfpPortStats) Collisions(b []byte) (uint64, error) {
	return packets.ReadB64(b, 96)
}
func (o OfpPortStats) SetCollisions(b []byte, vn uint64) error {
	return packets.WriteB64(b, 96, vn)
}
func (o OfpPortStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 104)
}
func (o OfpPortStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 104, vn)
}
func (o OfpPortStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 108)
}
func (o OfpPortStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 108, vn)
}
