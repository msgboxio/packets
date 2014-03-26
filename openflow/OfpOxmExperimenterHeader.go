package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpOxmExperimenterHeader = 8

type OfpOxmExperimenterHeader struct {
	// oxm_header uint32
	// experimenter uint32
}

func (o OfpOxmExperimenterHeader) OxmHeader(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpOxmExperimenterHeader) SetOxmHeader(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpOxmExperimenterHeader) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpOxmExperimenterHeader) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
