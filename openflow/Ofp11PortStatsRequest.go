package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11PortStatsRequest = 8

type Ofp11PortStatsRequest struct {
	// port_no uint32
	// pad [4]byte
}

func (o Ofp11PortStatsRequest) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11PortStatsRequest) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
