package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpExperimenterHeader = 16

type OfpExperimenterHeader struct {
	// header ofp_header
	// experimenter uint32
	// exp_type uint32
}

func (o OfpExperimenterHeader) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpExperimenterHeader) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpExperimenterHeader) ExpType(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpExperimenterHeader) SetExpType(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
