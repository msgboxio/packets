package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionSetQueue = 8

type Ofp11ActionSetQueue struct {
	// msg_type uint16
	// len uint16
	// queue_id uint32
}

func (o Ofp11ActionSetQueue) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionSetQueue) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionSetQueue) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionSetQueue) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionSetQueue) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11ActionSetQueue) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
