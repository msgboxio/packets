package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11QueueStats = 32

type Ofp11QueueStats struct {
	// port_no uint32
	// queue_id uint32
	// tx_bytes uint64
	// tx_packets uint64
	// tx_errors uint64
}

func (o Ofp11QueueStats) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11QueueStats) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11QueueStats) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11QueueStats) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11QueueStats) TxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp11QueueStats) SetTxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o Ofp11QueueStats) TxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o Ofp11QueueStats) SetTxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o Ofp11QueueStats) TxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o Ofp11QueueStats) SetTxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
