package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpInstructionExperimenter = 8

type OfpInstructionExperimenter struct {
	// msg_type uint16
	// len uint16
	// experimenter uint32
}

func (o OfpInstructionExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpInstructionExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpInstructionExperimenter) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpInstructionExperimenter) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpInstructionExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpInstructionExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
