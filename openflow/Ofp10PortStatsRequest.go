package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PortStatsRequest = 8

type Ofp10PortStatsRequest struct {
	// port_no uint16
	// pad [6]byte
}

func (o Ofp10PortStatsRequest) PortNo(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10PortStatsRequest) SetPortNo(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
