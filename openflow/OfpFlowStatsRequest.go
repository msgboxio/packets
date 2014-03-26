package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpFlowStatsRequest = 40

type OfpFlowStatsRequest struct {
	// table_id uint8
	// pad [3]byte
	// out_port uint32
	// out_group uint32
	// pad2 [4]byte
	// cookie uint64
	// cookie_mask uint64
	// match ofp_match
}

func (o OfpFlowStatsRequest) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o OfpFlowStatsRequest) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o OfpFlowStatsRequest) OutPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpFlowStatsRequest) SetOutPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpFlowStatsRequest) OutGroup(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpFlowStatsRequest) SetOutGroup(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpFlowStatsRequest) Pad2(b []byte) ([]byte, error) {
	return b[12:16], nil
}
func (o OfpFlowStatsRequest) SetPad2(b []byte, vn []byte) error {
	copy(b[12:16], vn)
	return nil
}
func (o OfpFlowStatsRequest) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpFlowStatsRequest) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o OfpFlowStatsRequest) CookieMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 24)
}
func (o OfpFlowStatsRequest) SetCookieMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 24, vn)
}
