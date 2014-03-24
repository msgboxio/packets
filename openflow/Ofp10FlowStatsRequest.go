package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10FlowStatsRequest = 44

type Ofp10FlowStatsRequest struct {
	// match ofp10_match
	// table_id uint8
	// pad uint8
	// out_port uint16
}

func (o Ofp10FlowStatsRequest) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 40)
}
func (o Ofp10FlowStatsRequest) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 40, vn)
}
func (o Ofp10FlowStatsRequest) OutPort(b []byte) (uint16, error) {
	return packets.ReadB16(b, 42)
}
func (o Ofp10FlowStatsRequest) SetOutPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 42, vn)
}
