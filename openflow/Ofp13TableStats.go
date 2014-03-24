package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableStats = 24

type Ofp13TableStats struct {
	// table_id uint8
	// pad [3]byte
	// active_count uint32
	// lookup_count uint64
	// matched_count uint64
}

func (o Ofp13TableStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o Ofp13TableStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o Ofp13TableStats) ActiveCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13TableStats) SetActiveCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp13TableStats) LookupCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp13TableStats) SetLookupCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o Ofp13TableStats) MatchedCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o Ofp13TableStats) SetMatchedCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
