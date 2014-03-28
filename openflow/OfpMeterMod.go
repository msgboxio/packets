package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpMeterMod = 16

type OfpMeterMod struct {
	// header ofp_header
	// command uint16
	// flags uint16
	// meter_id uint32
}

func (o OfpMeterMod) Command(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpMeterMod) SetCommand(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpMeterMod) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o OfpMeterMod) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
func (o OfpMeterMod) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpMeterMod) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
