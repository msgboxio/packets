package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PortMod = 24

type Ofp10PortMod struct {
	// port_no uint16
	// hw_addr [6]byte
	// config uint32
	// mask uint32
	// advertise uint32
	// pad [4]byte
}

func (o Ofp10PortMod) PortNo(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10PortMod) SetPortNo(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10PortMod) HwAddr(b []byte) ([]byte, error) {
	return b[2:8], nil
}
func (o Ofp10PortMod) SetHwAddr(b []byte, vn []byte) error {
	copy(b[2:8], vn)
	return nil
}
func (o Ofp10PortMod) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 3)
}
func (o Ofp10PortMod) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 3, vn)
}
func (o Ofp10PortMod) Mask(b []byte) (uint32, error) {
	return packets.ReadB32(b, 7)
}
func (o Ofp10PortMod) SetMask(b []byte, vn uint32) error {
	return packets.WriteB32(b, 7, vn)
}
func (o Ofp10PortMod) Advertise(b []byte) (uint32, error) {
	return packets.ReadB32(b, 11)
}
func (o Ofp10PortMod) SetAdvertise(b []byte, vn uint32) error {
	return packets.WriteB32(b, 11, vn)
}
