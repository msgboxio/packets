package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpQueueGetConfigReply = 16

type OfpQueueGetConfigReply struct {
	// header ofp_header
	// port uint32
	// pad [4]byte
}

func (o OfpQueueGetConfigReply) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpQueueGetConfigReply) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
