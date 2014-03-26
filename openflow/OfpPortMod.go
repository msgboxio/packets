package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpPortMod = 40

type OfpPortMod struct {
	// header ofp_header
	// port_no uint32
	// pad [4]byte
	// hw_addr [6]byte
	// pad2 [2]byte
	// config uint32
	// mask uint32
	// advertise uint32
	// pad3 [4]byte
}

func (o OfpPortMod) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpPortMod) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpPortMod) HwAddr(b []byte) ([]byte, error) {
	return b[16:22], nil
}
func (o OfpPortMod) SetHwAddr(b []byte, vn []byte) error {
	copy(b[16:22], vn)
	return nil
}
func (o OfpPortMod) Pad2(b []byte) ([]byte, error) {
	return b[22:24], nil
}
func (o OfpPortMod) SetPad2(b []byte, vn []byte) error {
	copy(b[22:24], vn)
	return nil
}
func (o OfpPortMod) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o OfpPortMod) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
func (o OfpPortMod) Mask(b []byte) (uint32, error) {
	return packets.ReadB32(b, 28)
}
func (o OfpPortMod) SetMask(b []byte, vn uint32) error {
	return packets.WriteB32(b, 28, vn)
}
func (o OfpPortMod) Advertise(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o OfpPortMod) SetAdvertise(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o OfpPortMod) Pad3(b []byte) ([]byte, error) {
	return b[36:40], nil
}
func (o OfpPortMod) SetPad3(b []byte, vn []byte) error {
	copy(b[36:40], vn)
	return nil
}
