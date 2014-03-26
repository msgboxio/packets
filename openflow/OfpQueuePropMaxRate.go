package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpQueuePropMaxRate = 16

type OfpQueuePropMaxRate struct {
	// prop_header ofp_queue_prop_header
	// rate uint16
	// pad [6]byte
}

func (o OfpQueuePropMaxRate) Rate(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpQueuePropMaxRate) SetRate(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
