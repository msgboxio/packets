package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11Match = 88

type Ofp11Match struct {
	// omh ofp11_match_header
	// in_port uint32
	// wildcards uint32
	// dl_src [6]byte
	// dl_src_mask [6]byte
	// dl_dst [6]byte
	// dl_dst_mask [6]byte
	// dl_vlan uint16
	// dl_vlan_pcp uint8
	// pad1 [1]byte
	// dl_type uint16
	// nw_tos uint8
	// nw_proto uint8
	// nw_src uint32
	// nw_src_mask uint32
	// nw_dst uint32
	// nw_dst_mask uint32
	// tp_src uint16
	// tp_dst uint16
	// mpls_label uint32
	// mpls_tc uint8
	// pad2 [3]byte
	// metadata uint64
	// metadata_mask uint64
}

func (o Ofp11Match) InPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11Match) SetInPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11Match) Wildcards(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp11Match) SetWildcards(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp11Match) DlSrc(b []byte) ([]byte, error) {
	return b[12:18], nil
}
func (o Ofp11Match) SetDlSrc(b []byte, vn []byte) error {
	copy(b[12:18], vn)
	return nil
}
func (o Ofp11Match) DlSrcMask(b []byte) ([]byte, error) {
	return b[13:19], nil
}
func (o Ofp11Match) SetDlSrcMask(b []byte, vn []byte) error {
	copy(b[13:19], vn)
	return nil
}
func (o Ofp11Match) DlDst(b []byte) ([]byte, error) {
	return b[14:20], nil
}
func (o Ofp11Match) SetDlDst(b []byte, vn []byte) error {
	copy(b[14:20], vn)
	return nil
}
func (o Ofp11Match) DlDstMask(b []byte) ([]byte, error) {
	return b[15:21], nil
}
func (o Ofp11Match) SetDlDstMask(b []byte, vn []byte) error {
	copy(b[15:21], vn)
	return nil
}
func (o Ofp11Match) DlVlan(b []byte) (uint16, error) {
	return packets.ReadB16(b, 16)
}
func (o Ofp11Match) SetDlVlan(b []byte, vn uint16) error {
	return packets.WriteB16(b, 16, vn)
}
func (o Ofp11Match) DlVlanPcp(b []byte) (uint8, error) {
	return packets.ReadB8(b, 18)
}
func (o Ofp11Match) SetDlVlanPcp(b []byte, vn uint8) error {
	return packets.WriteB8(b, 18, vn)
}
func (o Ofp11Match) Pad1(b []byte) ([]byte, error) {
	return b[19:20], nil
}
func (o Ofp11Match) SetPad1(b []byte, vn []byte) error {
	copy(b[19:20], vn)
	return nil
}
func (o Ofp11Match) DlType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 20)
}
func (o Ofp11Match) SetDlType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 20, vn)
}
func (o Ofp11Match) NwTos(b []byte) (uint8, error) {
	return packets.ReadB8(b, 22)
}
func (o Ofp11Match) SetNwTos(b []byte, vn uint8) error {
	return packets.WriteB8(b, 22, vn)
}
func (o Ofp11Match) NwProto(b []byte) (uint8, error) {
	return packets.ReadB8(b, 23)
}
func (o Ofp11Match) SetNwProto(b []byte, vn uint8) error {
	return packets.WriteB8(b, 23, vn)
}
func (o Ofp11Match) NwSrc(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o Ofp11Match) SetNwSrc(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
func (o Ofp11Match) NwSrcMask(b []byte) (uint32, error) {
	return packets.ReadB32(b, 28)
}
func (o Ofp11Match) SetNwSrcMask(b []byte, vn uint32) error {
	return packets.WriteB32(b, 28, vn)
}
func (o Ofp11Match) NwDst(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o Ofp11Match) SetNwDst(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o Ofp11Match) NwDstMask(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o Ofp11Match) SetNwDstMask(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
func (o Ofp11Match) TpSrc(b []byte) (uint16, error) {
	return packets.ReadB16(b, 40)
}
func (o Ofp11Match) SetTpSrc(b []byte, vn uint16) error {
	return packets.WriteB16(b, 40, vn)
}
func (o Ofp11Match) TpDst(b []byte) (uint16, error) {
	return packets.ReadB16(b, 42)
}
func (o Ofp11Match) SetTpDst(b []byte, vn uint16) error {
	return packets.WriteB16(b, 42, vn)
}
func (o Ofp11Match) MplsLabel(b []byte) (uint32, error) {
	return packets.ReadB32(b, 44)
}
func (o Ofp11Match) SetMplsLabel(b []byte, vn uint32) error {
	return packets.WriteB32(b, 44, vn)
}
func (o Ofp11Match) MplsTc(b []byte) (uint8, error) {
	return packets.ReadB8(b, 48)
}
func (o Ofp11Match) SetMplsTc(b []byte, vn uint8) error {
	return packets.WriteB8(b, 48, vn)
}
func (o Ofp11Match) Pad2(b []byte) ([]byte, error) {
	return b[49:52], nil
}
func (o Ofp11Match) SetPad2(b []byte, vn []byte) error {
	copy(b[49:52], vn)
	return nil
}
func (o Ofp11Match) Metadata(b []byte) (uint64, error) {
	return packets.ReadB64(b, 50)
}
func (o Ofp11Match) SetMetadata(b []byte, vn uint64) error {
	return packets.WriteB64(b, 50, vn)
}
func (o Ofp11Match) MetadataMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 58)
}
func (o Ofp11Match) SetMetadataMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 58, vn)
}
