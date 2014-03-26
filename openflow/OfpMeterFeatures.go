package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMeterFeatures = 16

type OfpMeterFeatures struct {
	// max_meter uint32
	// band_types uint32
	// capabilities uint32
	// max_bands uint8
	// max_color uint8
	// pad [2]byte
}

func (o OfpMeterFeatures) MaxMeter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpMeterFeatures) SetMaxMeter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpMeterFeatures) BandTypes(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpMeterFeatures) SetBandTypes(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpMeterFeatures) Capabilities(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpMeterFeatures) SetCapabilities(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpMeterFeatures) MaxBands(b []byte) (uint8, error) {
	return packets.ReadB8(b, 12)
}
func (o OfpMeterFeatures) SetMaxBands(b []byte, vn uint8) error {
	return packets.WriteB8(b, 12, vn)
}
func (o OfpMeterFeatures) MaxColor(b []byte) (uint8, error) {
	return packets.ReadB8(b, 13)
}
func (o OfpMeterFeatures) SetMaxColor(b []byte, vn uint8) error {
	return packets.WriteB8(b, 13, vn)
}
