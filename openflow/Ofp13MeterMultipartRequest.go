package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterMultipartRequest = 8

type Ofp13MeterMultipartRequest struct {
	// meter_id uint32
	// pad [4]byte
}

func (o Ofp13MeterMultipartRequest) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp13MeterMultipartRequest) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
