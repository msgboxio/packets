package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMeterConfig = 8

type OfpMeterConfig struct {
	// length uint16
	// flags uint16
	// meter_id uint32
}

func (o OfpMeterConfig) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpMeterConfig) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpMeterConfig) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpMeterConfig) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpMeterConfig) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpMeterConfig) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
