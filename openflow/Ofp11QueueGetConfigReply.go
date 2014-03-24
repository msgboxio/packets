package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11QueueGetConfigReply = 8

type Ofp11QueueGetConfigReply struct {
	// port uint32
	// pad [4]byte
}

func (o Ofp11QueueGetConfigReply) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11QueueGetConfigReply) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
