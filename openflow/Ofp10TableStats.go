package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10TableStats = 64

type Ofp10TableStats struct {
	// table_id uint8
	// pad [3]byte
	// name char[32]
	// wildcards uint32
	// max_entries uint32
	// active_count uint32
	// lookup_count uint64
	// matched_count uint64
}

func (o Ofp10TableStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o Ofp10TableStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o Ofp10TableStats) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 2, 32, '\000')
	return string(slice), err
}
func (o Ofp10TableStats) SetName(b []byte, vn string) error {
	copy(b[2:], vn)
	return nil
}
func (o Ofp10TableStats) Wildcards(b []byte) (uint32, error) {
	return packets.ReadB32(b, 3)
}
func (o Ofp10TableStats) SetWildcards(b []byte, vn uint32) error {
	return packets.WriteB32(b, 3, vn)
}
func (o Ofp10TableStats) MaxEntries(b []byte) (uint32, error) {
	return packets.ReadB32(b, 7)
}
func (o Ofp10TableStats) SetMaxEntries(b []byte, vn uint32) error {
	return packets.WriteB32(b, 7, vn)
}
func (o Ofp10TableStats) ActiveCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 11)
}
func (o Ofp10TableStats) SetActiveCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 11, vn)
}
func (o Ofp10TableStats) LookupCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 15)
}
func (o Ofp10TableStats) SetLookupCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 15, vn)
}
func (o Ofp10TableStats) MatchedCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 23)
}
func (o Ofp10TableStats) SetMatchedCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 23, vn)
}
