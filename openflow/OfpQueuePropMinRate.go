package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpQueuePropMinRate = 16

type OfpQueuePropMinRate struct {
	// prop_header ofp_queue_prop_header
	// rate uint16
	// pad [6]byte
}

func (o OfpQueuePropMinRate) Rate(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpQueuePropMinRate) SetRate(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
