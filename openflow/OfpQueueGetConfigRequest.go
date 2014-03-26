package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpQueueGetConfigRequest = 16

type OfpQueueGetConfigRequest struct {
	// header ofp_header
	// port uint32
	// pad [4]byte
}

func (o OfpQueueGetConfigRequest) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpQueueGetConfigRequest) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
