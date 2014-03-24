package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12TableStats = 128

type Ofp12TableStats struct {
	// table_id uint8
	// pad [7]byte
	// name char[32]
	// match uint64
	// wildcards uint64
	// write_actions uint32
	// apply_actions uint32
	// write_setfields uint64
	// apply_setfields uint64
	// metadata_match uint64
	// metadata_write uint64
	// instructions uint32
	// config uint32
	// max_entries uint32
	// active_count uint32
	// lookup_count uint64
	// matched_count uint64
}

func (o Ofp12TableStats) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o Ofp12TableStats) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o Ofp12TableStats) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 8, 32, '\000')
	return string(slice), err
}
func (o Ofp12TableStats) SetName(b []byte, vn string) error {
	copy(b[8:], vn)
	return nil
}
func (o Ofp12TableStats) Match(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o Ofp12TableStats) SetMatch(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
func (o Ofp12TableStats) Wildcards(b []byte) (uint64, error) {
	return packets.ReadB64(b, 48)
}
func (o Ofp12TableStats) SetWildcards(b []byte, vn uint64) error {
	return packets.WriteB64(b, 48, vn)
}
func (o Ofp12TableStats) WriteActions(b []byte) (uint32, error) {
	return packets.ReadB32(b, 56)
}
func (o Ofp12TableStats) SetWriteActions(b []byte, vn uint32) error {
	return packets.WriteB32(b, 56, vn)
}
func (o Ofp12TableStats) ApplyActions(b []byte) (uint32, error) {
	return packets.ReadB32(b, 60)
}
func (o Ofp12TableStats) SetApplyActions(b []byte, vn uint32) error {
	return packets.WriteB32(b, 60, vn)
}
func (o Ofp12TableStats) WriteSetfields(b []byte) (uint64, error) {
	return packets.ReadB64(b, 64)
}
func (o Ofp12TableStats) SetWriteSetfields(b []byte, vn uint64) error {
	return packets.WriteB64(b, 64, vn)
}
func (o Ofp12TableStats) ApplySetfields(b []byte) (uint64, error) {
	return packets.ReadB64(b, 72)
}
func (o Ofp12TableStats) SetApplySetfields(b []byte, vn uint64) error {
	return packets.WriteB64(b, 72, vn)
}
func (o Ofp12TableStats) MetadataMatch(b []byte) (uint64, error) {
	return packets.ReadB64(b, 80)
}
func (o Ofp12TableStats) SetMetadataMatch(b []byte, vn uint64) error {
	return packets.WriteB64(b, 80, vn)
}
func (o Ofp12TableStats) MetadataWrite(b []byte) (uint64, error) {
	return packets.ReadB64(b, 88)
}
func (o Ofp12TableStats) SetMetadataWrite(b []byte, vn uint64) error {
	return packets.WriteB64(b, 88, vn)
}
func (o Ofp12TableStats) Instructions(b []byte) (uint32, error) {
	return packets.ReadB32(b, 96)
}
func (o Ofp12TableStats) SetInstructions(b []byte, vn uint32) error {
	return packets.WriteB32(b, 96, vn)
}
func (o Ofp12TableStats) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 100)
}
func (o Ofp12TableStats) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 100, vn)
}
func (o Ofp12TableStats) MaxEntries(b []byte) (uint32, error) {
	return packets.ReadB32(b, 104)
}
func (o Ofp12TableStats) SetMaxEntries(b []byte, vn uint32) error {
	return packets.WriteB32(b, 104, vn)
}
func (o Ofp12TableStats) ActiveCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 108)
}
func (o Ofp12TableStats) SetActiveCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 108, vn)
}
func (o Ofp12TableStats) LookupCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 112)
}
func (o Ofp12TableStats) SetLookupCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 112, vn)
}
func (o Ofp12TableStats) MatchedCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 120)
}
func (o Ofp12TableStats) SetMatchedCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 120, vn)
}
