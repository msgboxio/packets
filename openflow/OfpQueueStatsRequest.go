package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpQueueStatsRequest = 8

type OfpQueueStatsRequest struct {
	// port_no uint32
	// queue_id uint32
}

func (o OfpQueueStatsRequest) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpQueueStatsRequest) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpQueueStatsRequest) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpQueueStatsRequest) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
