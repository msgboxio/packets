package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10QueueStats = 32

type Ofp10QueueStats struct {
	// port_no uint16
	// pad [2]byte
	// queue_id uint32
	// tx_bytes uint64
	// tx_packets uint64
	// tx_errors uint64
}

func (o Ofp10QueueStats) PortNo(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10QueueStats) SetPortNo(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10QueueStats) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 3)
}
func (o Ofp10QueueStats) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 3, vn)
}
func (o Ofp10QueueStats) TxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 7)
}
func (o Ofp10QueueStats) SetTxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 7, vn)
}
func (o Ofp10QueueStats) TxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 15)
}
func (o Ofp10QueueStats) SetTxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 15, vn)
}
func (o Ofp10QueueStats) TxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 23)
}
func (o Ofp10QueueStats) SetTxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 23, vn)
}
