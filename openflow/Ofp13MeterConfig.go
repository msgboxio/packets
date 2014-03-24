package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterConfig = 8

type Ofp13MeterConfig struct {
	// length uint16
	// flags uint16
	// meter_id uint32
}

func (o Ofp13MeterConfig) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13MeterConfig) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13MeterConfig) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13MeterConfig) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13MeterConfig) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterConfig) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
