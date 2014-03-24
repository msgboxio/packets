package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp14AsyncConfigPropExperimenter = 12

type Ofp14AsyncConfigPropExperimenter struct {
	// msg_type uint16
	// length uint16
	// experimenter uint32
	// exp_type uint32
}

func (o Ofp14AsyncConfigPropExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp14AsyncConfigPropExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp14AsyncConfigPropExperimenter) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp14AsyncConfigPropExperimenter) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp14AsyncConfigPropExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp14AsyncConfigPropExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp14AsyncConfigPropExperimenter) ExpType(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp14AsyncConfigPropExperimenter) SetExpType(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
