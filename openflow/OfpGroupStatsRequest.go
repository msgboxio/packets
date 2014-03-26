package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpGroupStatsRequest = 8

type OfpGroupStatsRequest struct {
	// group_id uint32
	// pad [4]byte
}

func (o OfpGroupStatsRequest) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpGroupStatsRequest) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
