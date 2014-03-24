package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11TableStats = 88

type Ofp11TableStats struct {
	// table_id uint8
	// pad [7]byte
	// name char[32]
	// wildcards uint32
	// match uint32
	// instructions uint32
	// write_actions uint32
	// apply_actions uint32
	// config uint32
	// max_entries uint32
	// active_count uint32
	// lookup_count uint64
	// matched_count uint64
}

func (o Ofp11TableStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o Ofp11TableStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o Ofp11TableStats) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 2, 32, '\000')
	return string(slice), err
}
func (o Ofp11TableStats) SetName(b []byte, vn string) error {
	copy(b[2:], vn)
	return nil
}
func (o Ofp11TableStats) Wildcards(b []byte) (uint32, error) {
	return packets.ReadB32(b, 3)
}
func (o Ofp11TableStats) SetWildcards(b []byte, vn uint32) error {
	return packets.WriteB32(b, 3, vn)
}
func (o Ofp11TableStats) Match(b []byte) (uint32, error) {
	return packets.ReadB32(b, 7)
}
func (o Ofp11TableStats) SetMatch(b []byte, vn uint32) error {
	return packets.WriteB32(b, 7, vn)
}
func (o Ofp11TableStats) Instructions(b []byte) (uint32, error) {
	return packets.ReadB32(b, 11)
}
func (o Ofp11TableStats) SetInstructions(b []byte, vn uint32) error {
	return packets.WriteB32(b, 11, vn)
}
func (o Ofp11TableStats) WriteActions(b []byte) (uint32, error) {
	return packets.ReadB32(b, 15)
}
func (o Ofp11TableStats) SetWriteActions(b []byte, vn uint32) error {
	return packets.WriteB32(b, 15, vn)
}
func (o Ofp11TableStats) ApplyActions(b []byte) (uint32, error) {
	return packets.ReadB32(b, 19)
}
func (o Ofp11TableStats) SetApplyActions(b []byte, vn uint32) error {
	return packets.WriteB32(b, 19, vn)
}
func (o Ofp11TableStats) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 23)
}
func (o Ofp11TableStats) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 23, vn)
}
func (o Ofp11TableStats) MaxEntries(b []byte) (uint32, error) {
	return packets.ReadB32(b, 27)
}
func (o Ofp11TableStats) SetMaxEntries(b []byte, vn uint32) error {
	return packets.WriteB32(b, 27, vn)
}
func (o Ofp11TableStats) ActiveCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 31)
}
func (o Ofp11TableStats) SetActiveCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 31, vn)
}
func (o Ofp11TableStats) LookupCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 35)
}
func (o Ofp11TableStats) SetLookupCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 35, vn)
}
func (o Ofp11TableStats) MatchedCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 43)
}
func (o Ofp11TableStats) SetMatchedCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 43, vn)
}
