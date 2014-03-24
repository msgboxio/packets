package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10QueueGetConfigReply = 8

type Ofp10QueueGetConfigReply struct {
	// port uint16
	// pad [6]byte
}

func (o Ofp10QueueGetConfigReply) Port(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10QueueGetConfigReply) SetPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
