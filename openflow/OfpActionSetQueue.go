package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionSetQueue = 8

type OfpActionSetQueue struct {
	// msg_type uint16
	// len uint16
	// queue_id uint32
}

func (o OfpActionSetQueue) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionSetQueue) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionSetQueue) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionSetQueue) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionSetQueue) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpActionSetQueue) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
