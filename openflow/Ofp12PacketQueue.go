package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12PacketQueue = 16

type Ofp12PacketQueue struct {
	// queue_id uint32
	// port uint32
	// len uint16
	// pad [6]byte
}

func (o Ofp12PacketQueue) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp12PacketQueue) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp12PacketQueue) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp12PacketQueue) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp12PacketQueue) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp12PacketQueue) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
