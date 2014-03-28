package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpPortStatus = 80

type OfpPortStatus struct {
	// header ofp_header
	// reason uint8
	// pad [7]byte
	// desc ofp_port
}

func (o OfpPortStatus) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 8)
}
func (o OfpPortStatus) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 8, vn)
}
