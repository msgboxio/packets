package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpExperimenterMultipartHeader = 8

type OfpExperimenterMultipartHeader struct {
	// experimenter uint32
	// exp_type uint32
}

func (o OfpExperimenterMultipartHeader) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpExperimenterMultipartHeader) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpExperimenterMultipartHeader) ExpType(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpExperimenterMultipartHeader) SetExpType(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
