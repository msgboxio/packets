package packets

const SizeofFlow = 168

func NewFlow() *Flow {
	flow := new(Flow)
	flow.b = make([]byte, SizeofFlow)
	return flow
}

type Flow struct {
	// tun_id uint64
	// tun_ip_src uint32
	// tun_ip_dst uint32
	// tun_flags uint16
	// tun_ip_tos uint8
	// tun_ip_ttl uint8
	// metadata uint64
	// regs [8]uint32
	// skb_priority uint32
	// pkt_mark uint32
	// in_port uint32
	// dl_src [6]byte
	// dl_dst [6]byte
	// dl_type uint16
	// vlan_tci uint16
	// mpls_lse uint32
	// ipv6_src [16]byte
	// ipv6_dst [16]byte
	// nd_target [16]byte
	// ipv6_label uint32
	// nw_frag uint8
	// nw_tos uint8
	// nw_ttl uint8
	// nw_proto uint8
	// arp_sha [6]byte
	// arp_tha [6]byte
	// tcp_flags uint16
	// pad uint16
	// tp_src uint16
	// tp_dst uint16
	b []byte
}

func (o *Flow) TunId() (v uint64) {
	v, _ = ReadB64(o.b, 0)
	return
}
func (o *Flow) SetTunId(vn uint64) error {
	return WriteB64(o.b, 0, vn)
}
func (o *Flow) TunIpSrc() (v uint32) {
	v, _ = ReadB32(o.b, 8)
	return
}
func (o *Flow) SetTunIpSrc(vn uint32) error {
	return WriteB32(o.b, 8, vn)
}
func (o *Flow) TunIpDst() (v uint32) {
	v, _ = ReadB32(o.b, 12)
	return
}
func (o *Flow) SetTunIpDst(vn uint32) error {
	return WriteB32(o.b, 12, vn)
}
func (o *Flow) TunFlags() (v uint16) {
	v, _ = ReadB16(o.b, 16)
	return
}
func (o *Flow) SetTunFlags(vn uint16) error {
	return WriteB16(o.b, 16, vn)
}
func (o *Flow) TunIpTos() (v uint8) {
	v, _ = ReadB8(o.b, 18)
	return
}
func (o *Flow) SetTunIpTos(vn uint8) error {
	return WriteB8(o.b, 18, vn)
}
func (o *Flow) TunIpTtl() (v uint8) {
	v, _ = ReadB8(o.b, 19)
	return
}
func (o *Flow) SetTunIpTtl(vn uint8) error {
	return WriteB8(o.b, 19, vn)
}
func (o *Flow) Metadata() (v uint64) {
	v, _ = ReadB64(o.b, 20)
	return
}
func (o *Flow) SetMetadata(vn uint64) error {
	return WriteB64(o.b, 20, vn)
}
func (o *Flow) Regs(ret []uint32) (err error) {
	for i := 0; i < 8; i++ {
		ret[i], err = ReadB32(o.b, 28+i*4)
	}
	return
}
func (o *Flow) SetRegs(vn []uint32) (err error) {
	for i := 0; i < 8; i++ {
		err = WriteB32(o.b, 28+i*4, vn[i])
	}
	return
}
func (o *Flow) SkbPriority() (v uint32) {
	v, _ = ReadB32(o.b, 60)
	return
}
func (o *Flow) SetSkbPriority(vn uint32) error {
	return WriteB32(o.b, 60, vn)
}
func (o *Flow) PktMark() (v uint32) {
	v, _ = ReadB32(o.b, 64)
	return
}
func (o *Flow) SetPktMark(vn uint32) error {
	return WriteB32(o.b, 64, vn)
}
func (o *Flow) InPort() (v uint32) {
	v, _ = ReadB32(o.b, 68)
	return
}
func (o *Flow) SetInPort(vn uint32) error {
	return WriteB32(o.b, 68, vn)
}
func (o *Flow) DlSrc() []byte {
	return o.b[72:78]
}
func (o *Flow) SetDlSrc(vn []byte) error {
	copy(o.b[72:78], vn)
	return nil
}
func (o *Flow) DlDst() []byte {
	return o.b[78:84]
}
func (o *Flow) SetDlDst(vn []byte) error {
	copy(o.b[78:84], vn)
	return nil
}
func (o *Flow) DlType() (v uint16) {
	v, _ = ReadB16(o.b, 84)
	return
}
func (o *Flow) SetDlType(vn uint16) error {
	return WriteB16(o.b, 84, vn)
}
func (o *Flow) VlanTci() (v uint16) {
	v, _ = ReadB16(o.b, 86)
	return
}
func (o *Flow) SetVlanTci(vn uint16) error {
	return WriteB16(o.b, 86, vn)
}
func (o *Flow) MplsLse() (v uint32) {
	v, _ = ReadB32(o.b, 88)
	return
}
func (o *Flow) SetMplsLse(vn uint32) error {
	return WriteB32(o.b, 88, vn)
}

func (o *Flow) NwSrc() (v uint32) {
	v, _ = ReadB32(o.b, 92)
	return
}
func (o *Flow) SetNwSrc(vn uint32) error {
	return WriteB32(o.b, 92, vn)
}
func (o *Flow) NwDst() (v uint32) {
	v, _ = ReadB32(o.b, 96)
	return
}
func (o *Flow) SetNwDst(vn uint32) error {
	return WriteB32(o.b, 96, vn)
}

func (o *Flow) Ipv6Src() []byte {
	return o.b[92:108]
}
func (o *Flow) SetIpv6Src(vn []byte) error {
	copy(o.b[92:108], vn)
	return nil
}
func (o *Flow) Ipv6Dst() []byte {
	return o.b[108:124]
}
func (o *Flow) SetIpv6Dst(vn []byte) error {
	copy(o.b[108:124], vn)
	return nil
}
func (o *Flow) NdTarget() []byte {
	return o.b[124:140]
}
func (o *Flow) SetNdTarget(vn []byte) error {
	copy(o.b[124:140], vn)
	return nil
}
func (o *Flow) Ipv6Label() (v uint32) {
	v, _ = ReadB32(o.b, 140)
	return
}
func (o *Flow) SetIpv6Label(vn uint32) error {
	return WriteB32(o.b, 140, vn)
}
func (o *Flow) NwFrag() (v uint8) {
	v, _ = ReadB8(o.b, 144)
	return
}
func (o *Flow) SetNwFrag(vn uint8) error {
	return WriteB8(o.b, 144, vn)
}
func (o *Flow) NwTos() (v uint8) {
	v, _ = ReadB8(o.b, 145)
	return
}
func (o *Flow) SetNwTos(vn uint8) error {
	return WriteB8(o.b, 145, vn)
}
func (o *Flow) NwTtl() (v uint8) {
	v, _ = ReadB8(o.b, 146)
	return
}
func (o *Flow) SetNwTtl(vn uint8) error {
	return WriteB8(o.b, 146, vn)
}
func (o *Flow) NwProto() (v uint8) {
	v, _ = ReadB8(o.b, 147)
	return
}
func (o *Flow) SetNwProto(vn uint8) error {
	return WriteB8(o.b, 147, vn)
}
func (o *Flow) ArpSha() []byte {
	return o.b[148:154]
}
func (o *Flow) SetArpSha(vn []byte) error {
	copy(o.b[148:154], vn)
	return nil
}
func (o *Flow) ArpTha() []byte {
	return o.b[154:160]
}
func (o *Flow) SetArpTha(vn []byte) error {
	copy(o.b[154:160], vn)
	return nil
}
func (o *Flow) TcpFlags() (v uint16) {
	v, _ = ReadB16(o.b, 160)
	return
}
func (o *Flow) SetTcpFlags(vn uint16) error {
	return WriteB16(o.b, 160, vn)
}
func (o *Flow) TpSrc() (v uint16) {
	v, _ = ReadB16(o.b, 164)
	return
}
func (o *Flow) SetTpSrc(vn uint16) error {
	return WriteB16(o.b, 164, vn)
}
func (o *Flow) TpDst() (v uint16) {
	v, _ = ReadB16(o.b, 166)
	return
}
func (o *Flow) SetTpDst(vn uint16) error {
	return WriteB16(o.b, 166, vn)
}
