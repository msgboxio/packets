package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11PortMod = 32

type Ofp11PortMod struct {
	// port_no uint32
	// pad [4]byte
	// hw_addr [6]byte
	// pad2 [2]byte
	// config uint32
	// mask uint32
	// advertise uint32
	// pad3 [4]byte
}

func (o Ofp11PortMod) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11PortMod) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11PortMod) HwAddr(b []byte) ([]byte, error) {
	return b[5:11], nil
}
func (o Ofp11PortMod) SetHwAddr(b []byte, vn []byte) error {
	copy(b[5:11], vn)
	return nil
}
func (o Ofp11PortMod) Pad2(b []byte) ([]byte, error) {
	return b[6:8], nil
}
func (o Ofp11PortMod) SetPad2(b []byte, vn []byte) error {
	copy(b[6:8], vn)
	return nil
}
func (o Ofp11PortMod) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 7)
}
func (o Ofp11PortMod) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 7, vn)
}
func (o Ofp11PortMod) Mask(b []byte) (uint32, error) {
	return packets.ReadB32(b, 11)
}
func (o Ofp11PortMod) SetMask(b []byte, vn uint32) error {
	return packets.WriteB32(b, 11, vn)
}
func (o Ofp11PortMod) Advertise(b []byte) (uint32, error) {
	return packets.ReadB32(b, 15)
}
func (o Ofp11PortMod) SetAdvertise(b []byte, vn uint32) error {
	return packets.WriteB32(b, 15, vn)
}
func (o Ofp11PortMod) Pad3(b []byte) ([]byte, error) {
	return b[19:23], nil
}
func (o Ofp11PortMod) SetPad3(b []byte, vn []byte) error {
	copy(b[19:23], vn)
	return nil
}
