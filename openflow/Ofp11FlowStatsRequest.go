package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11FlowStatsRequest = 32

type Ofp11FlowStatsRequest struct {
	// table_id uint8
	// pad [3]byte
	// out_port uint32
	// out_group uint32
	// pad2 [4]byte
	// cookie uint64
	// cookie_mask uint64
}

func (o Ofp11FlowStatsRequest) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o Ofp11FlowStatsRequest) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o Ofp11FlowStatsRequest) OutPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 2)
}
func (o Ofp11FlowStatsRequest) SetOutPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 2, vn)
}
func (o Ofp11FlowStatsRequest) OutGroup(b []byte) (uint32, error) {
	return packets.ReadB32(b, 6)
}
func (o Ofp11FlowStatsRequest) SetOutGroup(b []byte, vn uint32) error {
	return packets.WriteB32(b, 6, vn)
}
func (o Ofp11FlowStatsRequest) Pad2(b []byte) ([]byte, error) {
	return b[10:14], nil
}
func (o Ofp11FlowStatsRequest) SetPad2(b []byte, vn []byte) error {
	copy(b[10:14], vn)
	return nil
}
func (o Ofp11FlowStatsRequest) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 11)
}
func (o Ofp11FlowStatsRequest) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 11, vn)
}
func (o Ofp11FlowStatsRequest) CookieMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 19)
}
func (o Ofp11FlowStatsRequest) SetCookieMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 19, vn)
}
