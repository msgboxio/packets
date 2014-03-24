package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PhyPort = 48

type Ofp10PhyPort struct {
	// port_no uint16
	// hw_addr [6]byte
	// name char[16]
	// config uint32
	// state uint32
	// curr uint32
	// advertised uint32
	// supported uint32
	// peer uint32
}

func (o Ofp10PhyPort) PortNo(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10PhyPort) SetPortNo(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10PhyPort) HwAddr(b []byte) ([]byte, error) {
	return b[2:8], nil
}
func (o Ofp10PhyPort) SetHwAddr(b []byte, vn []byte) error {
	copy(b[2:8], vn)
	return nil
}
func (o Ofp10PhyPort) Name(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 3, 16, '\000')
	return string(slice), err
}
func (o Ofp10PhyPort) SetName(b []byte, vn string) error {
	copy(b[3:], vn)
	return nil
}
func (o Ofp10PhyPort) Config(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp10PhyPort) SetConfig(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp10PhyPort) State(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp10PhyPort) SetState(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp10PhyPort) Curr(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp10PhyPort) SetCurr(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o Ofp10PhyPort) Advertised(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o Ofp10PhyPort) SetAdvertised(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
func (o Ofp10PhyPort) Supported(b []byte) (uint32, error) {
	return packets.ReadB32(b, 20)
}
func (o Ofp10PhyPort) SetSupported(b []byte, vn uint32) error {
	return packets.WriteB32(b, 20, vn)
}
func (o Ofp10PhyPort) Peer(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o Ofp10PhyPort) SetPeer(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
