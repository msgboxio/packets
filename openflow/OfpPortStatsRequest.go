package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpPortStatsRequest = 8

type OfpPortStatsRequest struct {
	// port_no uint32
	// pad [4]byte
}

func (o OfpPortStatsRequest) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpPortStatsRequest) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
