package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableFeatures = 64

type Ofp13TableFeatures struct {
	// length uint16
	// table_id uint8
	// pad [5]byte
	// name char[32]
	// metadata_match uint64
	// metadata_write uint64
	// config uint32
	// max_entries uint32
}

func (o Ofp13TableFeatures) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13TableFeatures) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13TableFeatures) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 2)
}
func (o Ofp13TableFeatures) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 2, vn)
}
func (o Ofp13TableFeatures) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 8, 32, '\000')
	return string(slice), err
}
func (o Ofp13TableFeatures) SetName(b []byte, vn string) error {
	copy(b[8:], vn)
	return nil
}
func (o Ofp13TableFeatures) MetadataMatch(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o Ofp13TableFeatures) SetMetadataMatch(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
func (o Ofp13TableFeatures) MetadataWrite(b []byte) (uint64, error) {
	return packets.ReadB64(b, 48)
}
func (o Ofp13TableFeatures) SetMetadataWrite(b []byte, vn uint64) error {
	return packets.WriteB64(b, 48, vn)
}
func (o Ofp13TableFeatures) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 56)
}
func (o Ofp13TableFeatures) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 56, vn)
}
func (o Ofp13TableFeatures) MaxEntries(b []byte) (uint32, error) {
	return packets.ReadB32(b, 60)
}
func (o Ofp13TableFeatures) SetMaxEntries(b []byte, vn uint32) error {
	return packets.WriteB32(b, 60, vn)
}
