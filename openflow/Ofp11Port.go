package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11Port = 64

type Ofp11Port struct {
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

func (o Ofp11Port) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11Port) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11Port) HwAddr(b []byte) ([]byte, error) {
	return b[5:11], nil
}
func (o Ofp11Port) SetHwAddr(b []byte, vn []byte) error {
	copy(b[5:11], vn)
	return nil
}
func (o Ofp11Port) Pad2(b []byte) ([]byte, error) {
	return b[6:8], nil
}
func (o Ofp11Port) SetPad2(b []byte, vn []byte) error {
	copy(b[6:8], vn)
	return nil
}
func (o Ofp11Port) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 7, 16, '\000')
	return string(slice), err
}
func (o Ofp11Port) SetName(b []byte, vn string) error {
	copy(b[7:], vn)
	return nil
}
func (o Ofp11Port) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp11Port) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp11Port) State(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp11Port) SetState(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o Ofp11Port) Curr(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o Ofp11Port) SetCurr(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
func (o Ofp11Port) Advertised(b []byte) (uint32, error) {
	return packets.ReadB32(b, 20)
}
func (o Ofp11Port) SetAdvertised(b []byte, vn uint32) error {
	return packets.WriteB32(b, 20, vn)
}
func (o Ofp11Port) Supported(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o Ofp11Port) SetSupported(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
func (o Ofp11Port) Peer(b []byte) (uint32, error) {
	return packets.ReadB32(b, 28)
}
func (o Ofp11Port) SetPeer(b []byte, vn uint32) error {
	return packets.WriteB32(b, 28, vn)
}
func (o Ofp11Port) CurrSpeed(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o Ofp11Port) SetCurrSpeed(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o Ofp11Port) MaxSpeed(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o Ofp11Port) SetMaxSpeed(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
