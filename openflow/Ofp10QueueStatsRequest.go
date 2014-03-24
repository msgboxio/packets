package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10QueueStatsRequest = 8

type Ofp10QueueStatsRequest struct {
	// port_no uint16
	// pad [2]byte
	// queue_id uint32
}

func (o Ofp10QueueStatsRequest) PortNo(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10QueueStatsRequest) SetPortNo(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10QueueStatsRequest) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 3)
}
func (o Ofp10QueueStatsRequest) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 3, vn)
}
