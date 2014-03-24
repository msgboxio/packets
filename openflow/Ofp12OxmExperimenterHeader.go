package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12OxmExperimenterHeader = 8

type Ofp12OxmExperimenterHeader struct {
	// oxm_header uint32
	// experimenter uint32
}

func (o Ofp12OxmExperimenterHeader) OxmHeader(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp12OxmExperimenterHeader) SetOxmHeader(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp12OxmExperimenterHeader) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp12OxmExperimenterHeader) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
