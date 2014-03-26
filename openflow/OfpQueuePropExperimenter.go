package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpQueuePropExperimenter = 16

type OfpQueuePropExperimenter struct {
	// prop_header ofp_queue_prop_header
	// experimenter uint32
	// pad [4]byte
	// data [0]byte
}

func (o OfpQueuePropExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpQueuePropExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpQueuePropExperimenter) Data(b []byte) ([]byte, error) {
	return b[16:16], nil
}
func (o OfpQueuePropExperimenter) SetData(b []byte, vn []byte) error {
	copy(b[16:16], vn)
	return nil
}
