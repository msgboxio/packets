package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11InstructionExperimenter = 8

type Ofp11InstructionExperimenter struct {
	// msg_type uint16
	// len uint16
	// experimenter uint32
}

func (o Ofp11InstructionExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11InstructionExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11InstructionExperimenter) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11InstructionExperimenter) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11InstructionExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11InstructionExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
