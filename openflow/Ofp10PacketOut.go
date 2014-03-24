package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PacketOut = 8

type Ofp10PacketOut struct {
	// buffer_id uint32
	// in_port uint16
	// actions_len uint16
}

func (o Ofp10PacketOut) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp10PacketOut) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp10PacketOut) InPort(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp10PacketOut) SetInPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp10PacketOut) ActionsLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 6)
}
func (o Ofp10PacketOut) SetActionsLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 6, vn)
}
