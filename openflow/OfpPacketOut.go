package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpPacketOut = 24

type OfpPacketOut struct {
	// header ofp_header
	// buffer_id uint32
	// in_port uint32
	// actions_len uint16
	// pad [6]byte
}

func (o OfpPacketOut) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpPacketOut) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpPacketOut) InPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpPacketOut) SetInPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o OfpPacketOut) ActionsLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 16)
}
func (o OfpPacketOut) SetActionsLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 16, vn)
}
