package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpTableMod = 16

type OfpTableMod struct {
	// header ofp_header
	// table_id uint8
	// pad [3]byte
	// config uint32
}

func (o OfpTableMod) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 8)
}
func (o OfpTableMod) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 8, vn)
}
func (o OfpTableMod) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpTableMod) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
