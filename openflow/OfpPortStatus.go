package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpPortStatus = 8

type OfpPortStatus struct {
	// reason uint8
	// pad [7]byte
}

func (o OfpPortStatus) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o OfpPortStatus) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
