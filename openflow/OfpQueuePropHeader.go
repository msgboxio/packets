package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpQueuePropHeader = 8

type OfpQueuePropHeader struct {
	// property uint16
	// len uint16
	// pad [4]byte
}

func (o OfpQueuePropHeader) Property(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpQueuePropHeader) SetProperty(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpQueuePropHeader) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpQueuePropHeader) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
