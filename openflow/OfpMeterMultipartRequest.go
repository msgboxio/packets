package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpMeterMultipartRequest = 8

type OfpMeterMultipartRequest struct {
	// meter_id uint32
	// pad [4]byte
}

func (o OfpMeterMultipartRequest) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpMeterMultipartRequest) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
