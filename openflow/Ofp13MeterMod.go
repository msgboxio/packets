package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterMod = 8

type Ofp13MeterMod struct {
	// command uint16
	// flags uint16
	// meter_id uint32
}

func (o Ofp13MeterMod) Command(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13MeterMod) SetCommand(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13MeterMod) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13MeterMod) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13MeterMod) MeterId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterMod) SetMeterId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
