package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11PacketOut = 16

type Ofp11PacketOut struct {
	// buffer_id uint32
	// in_port uint32
	// actions_len uint16
	// pad [6]byte
}

func (o Ofp11PacketOut) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11PacketOut) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11PacketOut) InPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11PacketOut) SetInPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11PacketOut) ActionsLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp11PacketOut) SetActionsLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
