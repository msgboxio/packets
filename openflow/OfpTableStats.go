package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpTableStats = 24

type OfpTableStats struct {
	// table_id uint8
	// pad [3]byte
	// active_count uint32
	// lookup_count uint64
	// matched_count uint64
}

func (o OfpTableStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o OfpTableStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o OfpTableStats) ActiveCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpTableStats) SetActiveCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpTableStats) LookupCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpTableStats) SetLookupCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpTableStats) MatchedCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpTableStats) SetMatchedCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
