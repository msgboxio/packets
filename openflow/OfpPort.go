package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpPort = 64

type OfpPort struct {
	// port_no uint32
	// pad [4]byte
	// hw_addr [6]byte
	// pad2 [2]byte
	// name char[16]
	// config uint32
	// state uint32
	// curr uint32
	// advertised uint32
	// supported uint32
	// peer uint32
	// curr_speed uint32
	// max_speed uint32
}

func (o OfpPort) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpPort) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpPort) HwAddr(b []byte) ([]byte, error) {
	return b[8:14], nil
}
func (o OfpPort) SetHwAddr(b []byte, vn []byte) error {
	copy(b[8:14], vn)
	return nil
}
func (o OfpPort) Pad2(b []byte) ([]byte, error) {
	return b[14:16], nil
}
func (o OfpPort) SetPad2(b []byte, vn []byte) error {
	copy(b[14:16], vn)
	return nil
}
func (o OfpPort) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 16, 16, '\000')
	return string(slice), err
}
func (o OfpPort) SetName(b []byte, vn string) error {
	copy(b[16:], vn)
	return nil
}
func (o OfpPort) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o OfpPort) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o OfpPort) State(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o OfpPort) SetState(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
func (o OfpPort) Curr(b []byte) (uint32, error) {
	return packets.ReadB32(b, 40)
}
func (o OfpPort) SetCurr(b []byte, vn uint32) error {
	return packets.WriteB32(b, 40, vn)
}
func (o OfpPort) Advertised(b []byte) (uint32, error) {
	return packets.ReadB32(b, 44)
}
func (o OfpPort) SetAdvertised(b []byte, vn uint32) error {
	return packets.WriteB32(b, 44, vn)
}
func (o OfpPort) Supported(b []byte) (uint32, error) {
	return packets.ReadB32(b, 48)
}
func (o OfpPort) SetSupported(b []byte, vn uint32) error {
	return packets.WriteB32(b, 48, vn)
}
func (o OfpPort) Peer(b []byte) (uint32, error) {
	return packets.ReadB32(b, 52)
}
func (o OfpPort) SetPeer(b []byte, vn uint32) error {
	return packets.WriteB32(b, 52, vn)
}
func (o OfpPort) CurrSpeed(b []byte) (uint32, error) {
	return packets.ReadB32(b, 56)
}
func (o OfpPort) SetCurrSpeed(b []byte, vn uint32) error {
	return packets.WriteB32(b, 56, vn)
}
func (o OfpPort) MaxSpeed(b []byte) (uint32, error) {
	return packets.ReadB32(b, 60)
}
func (o OfpPort) SetMaxSpeed(b []byte, vn uint32) error {
	return packets.WriteB32(b, 60, vn)
}
