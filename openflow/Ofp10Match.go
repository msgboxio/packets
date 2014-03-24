package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10Match = 40

type Ofp10Match struct {
	// wildcards uint32
	// in_port uint16
	// dl_src [6]byte
	// dl_dst [6]byte
	// dl_vlan uint16
	// dl_vlan_pcp uint8
	// pad1 [1]byte
	// dl_type uint16
	// nw_tos uint8
	// nw_proto uint8
	// pad2 [2]byte
	// nw_src uint32
	// nw_dst uint32
	// tp_src uint16
	// tp_dst uint16
}

func (o Ofp10Match) Wildcards(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp10Match) SetWildcards(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp10Match) InPort(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp10Match) SetInPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp10Match) DlSrc(b []byte) ([]byte, error) {
	return b[6:12], nil
}
func (o Ofp10Match) SetDlSrc(b []byte, vn []byte) error {
	copy(b[6:12], vn)
	return nil
}
func (o Ofp10Match) DlDst(b []byte) ([]byte, error) {
	return b[7:13], nil
}
func (o Ofp10Match) SetDlDst(b []byte, vn []byte) error {
	copy(b[7:13], vn)
	return nil
}
func (o Ofp10Match) DlVlan(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp10Match) SetDlVlan(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o Ofp10Match) DlVlanPcp(b []byte) (uint8, error) {
	return packets.ReadB8(b, 10)
}
func (o Ofp10Match) SetDlVlanPcp(b []byte, vn uint8) error {
	return packets.WriteB8(b, 10, vn)
}
func (o Ofp10Match) Pad1(b []byte) ([]byte, error) {
	return b[11:12], nil
}
func (o Ofp10Match) SetPad1(b []byte, vn []byte) error {
	copy(b[11:12], vn)
	return nil
}
func (o Ofp10Match) DlType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 12)
}
func (o Ofp10Match) SetDlType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 12, vn)
}
func (o Ofp10Match) NwTos(b []byte) (uint8, error) {
	return packets.ReadB8(b, 14)
}
func (o Ofp10Match) SetNwTos(b []byte, vn uint8) error {
	return packets.WriteB8(b, 14, vn)
}
func (o Ofp10Match) NwProto(b []byte) (uint8, error) {
	return packets.ReadB8(b, 15)
}
func (o Ofp10Match) SetNwProto(b []byte, vn uint8) error {
	return packets.WriteB8(b, 15, vn)
}
func (o Ofp10Match) Pad2(b []byte) ([]byte, error) {
	return b[16:18], nil
}
func (o Ofp10Match) SetPad2(b []byte, vn []byte) error {
	copy(b[16:18], vn)
	return nil
}
func (o Ofp10Match) NwSrc(b []byte) (uint32, error) {
	return packets.ReadB32(b, 17)
}
func (o Ofp10Match) SetNwSrc(b []byte, vn uint32) error {
	return packets.WriteB32(b, 17, vn)
}
func (o Ofp10Match) NwDst(b []byte) (uint32, error) {
	return packets.ReadB32(b, 21)
}
func (o Ofp10Match) SetNwDst(b []byte, vn uint32) error {
	return packets.WriteB32(b, 21, vn)
}
func (o Ofp10Match) TpSrc(b []byte) (uint16, error) {
	return packets.ReadB16(b, 25)
}
func (o Ofp10Match) SetTpSrc(b []byte, vn uint16) error {
	return packets.WriteB16(b, 25, vn)
}
func (o Ofp10Match) TpDst(b []byte) (uint16, error) {
	return packets.ReadB16(b, 27)
}
func (o Ofp10Match) SetTpDst(b []byte, vn uint16) error {
	return packets.WriteB16(b, 27, vn)
}
