package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11TableMod = 8

type Ofp11TableMod struct {
	// table_id uint8
	// pad [3]byte
	// config uint32
}

func (o Ofp11TableMod) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o Ofp11TableMod) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o Ofp11TableMod) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 2)
}
func (o Ofp11TableMod) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 2, vn)
}
