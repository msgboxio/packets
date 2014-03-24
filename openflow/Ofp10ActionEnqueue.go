package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10ActionEnqueue = 16

type Ofp10ActionEnqueue struct {
	// msg_type uint16
	// len uint16
	// port uint16
	// pad [6]byte
	// queue_id uint32
}

func (o Ofp10ActionEnqueue) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10ActionEnqueue) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10ActionEnqueue) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp10ActionEnqueue) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp10ActionEnqueue) Port(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp10ActionEnqueue) SetPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp10ActionEnqueue) QueueId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 7)
}
func (o Ofp10ActionEnqueue) SetQueueId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 7, vn)
}
