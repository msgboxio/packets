package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterFeatures = 16

type Ofp13MeterFeatures struct {
	// max_meter uint32
	// band_types uint32
	// capabilities uint32
	// max_bands uint8
	// max_color uint8
	// pad [2]byte
}

func (o Ofp13MeterFeatures) MaxMeter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp13MeterFeatures) SetMaxMeter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp13MeterFeatures) BandTypes(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterFeatures) SetBandTypes(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp13MeterFeatures) Capabilities(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp13MeterFeatures) SetCapabilities(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp13MeterFeatures) MaxBands(b []byte) (uint8, error) {
	return packets.ReadB8(b, 12)
}
func (o Ofp13MeterFeatures) SetMaxBands(b []byte, vn uint8) error {
	return packets.WriteB8(b, 12, vn)
}
func (o Ofp13MeterFeatures) MaxColor(b []byte) (uint8, error) {
	return packets.ReadB8(b, 13)
}
func (o Ofp13MeterFeatures) SetMaxColor(b []byte, vn uint8) error {
	return packets.WriteB8(b, 13, vn)
}
