package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11QueueStatsRequest = 8

type Ofp11QueueStatsRequest struct {
	// port_no uint32
	// queue_id uint32
}

func (o Ofp11QueueStatsRequest) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11QueueStatsRequest) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11QueueStatsRequest) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11QueueStatsRequest) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
