package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12ExperimenterStatsHeader = 8

type Ofp12ExperimenterStatsHeader struct {
	// experimenter uint32
	// exp_type uint32
}

func (o Ofp12ExperimenterStatsHeader) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp12ExperimenterStatsHeader) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp12ExperimenterStatsHeader) ExpType(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp12ExperimenterStatsHeader) SetExpType(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
