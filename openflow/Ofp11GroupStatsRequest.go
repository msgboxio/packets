package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11GroupStatsRequest = 8

type Ofp11GroupStatsRequest struct {
	// group_id uint32
	// pad [4]byte
}

func (o Ofp11GroupStatsRequest) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11GroupStatsRequest) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
