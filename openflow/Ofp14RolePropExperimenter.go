package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp14RolePropExperimenter = 12

type Ofp14RolePropExperimenter struct {
	// msg_type uint16
	// length uint16
	// experimenter uint32
	// exp_type uint32
}

func (o Ofp14RolePropExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp14RolePropExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp14RolePropExperimenter) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp14RolePropExperimenter) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp14RolePropExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp14RolePropExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp14RolePropExperimenter) ExpType(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp14RolePropExperimenter) SetExpType(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
