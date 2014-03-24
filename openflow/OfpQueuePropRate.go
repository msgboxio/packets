package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpQueuePropRate = 16

type OfpQueuePropRate struct {
	// prop_header ofp_queue_prop_header
	// rate uint16
	// pad [6]byte
}

func (o OfpQueuePropRate) Rate(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpQueuePropRate) SetRate(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
