package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpQueueStats = 40

type OfpQueueStats struct {
	// port_no uint32
	// queue_id uint32
	// tx_bytes uint64
	// tx_packets uint64
	// tx_errors uint64
	// duration_sec uint32
	// duration_nsec uint32
}

func (o OfpQueueStats) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpQueueStats) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpQueueStats) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpQueueStats) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpQueueStats) TxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpQueueStats) SetTxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpQueueStats) TxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpQueueStats) SetTxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o OfpQueueStats) TxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o OfpQueueStats) SetTxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
func (o OfpQueueStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o OfpQueueStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o OfpQueueStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o OfpQueueStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
