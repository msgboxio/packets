package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PacketQueue = 8

type Ofp10PacketQueue struct {
	// queue_id uint32
	// len uint16
	// pad [2]byte
}

func (o Ofp10PacketQueue) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp10PacketQueue) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp10PacketQueue) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp10PacketQueue) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
