package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10QueueGetConfigRequest = 4

type Ofp10QueueGetConfigRequest struct {
	// port uint16
	// pad [2]byte
}

func (o Ofp10QueueGetConfigRequest) Port(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10QueueGetConfigRequest) SetPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
