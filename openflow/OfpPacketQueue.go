package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpPacketQueue = 16

type OfpPacketQueue struct {
	// queue_id uint32
	// port uint32
	// len uint16
	// pad [6]byte
}

func (o OfpPacketQueue) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpPacketQueue) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpPacketQueue) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpPacketQueue) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpPacketQueue) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpPacketQueue) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
