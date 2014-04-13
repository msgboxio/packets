package packets

import (
	"math"
)

type Match struct {
	Flow *Flow
	Mask *Flow
}

func NewMatch() *Match {
	match := new(Match)
	match.Flow = NewFlow()
	match.Mask = NewFlow()
	return match
}

func NewMatchFromFlow(flow *Flow) *Match {
	match := new(Match)
	match.Flow = flow
	match.Mask = NewFlow()
	match.SetMaskFromFlow(flow)
	return match
}

func fill(b []byte, v uint8) {
	for x := range b {
		b[x] = v
	}
}

/* Converts a flow into a match.  It sets the wildcard masks based on
 * the packet contents.  It will not set the mask for fields that do not
 * make sense for the packet type. */
func (match *Match) SetMaskFromFlow(flow *Flow) {
	match.Mask.SetDlType(0xffff)
	if flow.NwProto() != 0 {
		match.Mask.SetNwProto(0xff)
	}
	if flow.SkbPriority() != 0 {
		match.Mask.SetSkbPriority(math.MaxUint32)
	}
	if flow.PktMark() != 0 {
		match.Mask.SetPktMark(math.MaxUint32)
	}
	// for i := 0; i < FLOW_N_REGS; i++ {
	// 	if flow.u32_regs[i] != 0 {
	// 		match.Mask.regs[1] = math.MaxUint32
	// 	}
	// }
	if flow.TunIpDst() != 0 {
		if (flow.TunFlags() & FLOW_TNL_F_KEY) != 0 {
			match.Mask.SetTunId(math.MaxUint64)
		}
		match.Mask.SetTunIpSrc(math.MaxUint32)
		match.Mask.SetTunIpDst(math.MaxUint32)
		match.Mask.SetTunFlags(math.MaxUint16)
		match.Mask.SetTunIpTos(math.MaxUint8)
		match.Mask.SetTunIpTtl(math.MaxUint8)
	} else if flow.TunId() != 0 {
		match.Mask.SetTunId(math.MaxUint64)
	}
	match.Mask.SetMetadata(math.MaxUint64)
	match.Mask.SetInPort(math.MaxUint16)
	match.Mask.SetVlanTci(math.MaxUint16)
	fill(match.Mask.DlSrc(), math.MaxUint8)
	fill(match.Mask.DlDst(), math.MaxUint8)

	dl := flow.DlType()
	if dl == ETH_TYPE_IPV6 {
		fill(match.Mask.Ipv6Src(), math.MaxUint8)
		fill(match.Mask.Ipv6Dst(), math.MaxUint8)
		match.Mask.SetIpv6Label(math.MaxUint32)
	} else if dl == ETH_TYPE_IP || dl == ETH_TYPE_ARP || dl == ETH_TYPE_RARP {
		match.Mask.SetNwSrc(math.MaxUint32)
		match.Mask.SetNwDst(math.MaxUint32)
	} else if dl == ETH_TYPE_MPLS || dl == ETH_TYPE_MPLS_MCAST {
		match.Mask.SetMplsLse(math.MaxUint32)
	}
	if dl == ETH_TYPE_ARP || dl == ETH_TYPE_RARP {
		fill(match.Mask.ArpSha(), math.MaxUint8)
		fill(match.Mask.ArpTha(), math.MaxUint8)
	}
	if dl == ETH_TYPE_IP || dl == ETH_TYPE_IPV6 {
		match.Mask.SetNwTos(math.MaxUint8)
		match.Mask.SetNwTtl(math.MaxUint8)
		if flow.NwFrag() != 0 {
			match.Mask.SetNwFrag(math.MaxUint8)
			if (flow.NwFrag() & FLOW_NW_FRAG_LATER) != 0 {
				/* No transport layer header in later fragments. */
				return
			}
		}
		nw := flow.NwProto()
		if nw == IPPROTO_ICMP || nw == IPPROTO_ICMPV6 ||
			(flow.TpSrc() != 0 || flow.TpDst() != 0) {
			match.Mask.SetTpSrc(math.MaxUint16)
			match.Mask.SetTpDst(math.MaxUint16)
		}
		if nw == IPPROTO_TCP {
			match.Mask.SetTcpFlags(math.MaxUint16)
		}
		if nw == IPPROTO_ICMPV6 {
			fill(match.Mask.ArpSha(), math.MaxUint8)
			fill(match.Mask.ArpTha(), math.MaxUint8)
		}
	}
}

func (match *Match) SetInPort(port uint32) {
	match.Flow.SetInPort(port)
	match.Mask.SetInPort(math.MaxUint32)
}

func (match *Match) SetVlanPcp(pcp uint16) {
	pcp &= 0x07
	otci := match.Flow.VlanTci()
	match.Flow.SetVlanTci((otci & (^VLAN_PCP_MASK & 0xffff)) | (pcp << VLAN_PCP_SHIFT) | VLAN_CFI)
}

func (match *Match) SetVlanVid(vid uint16) {
	match.SetVlanVidMasked(vid, VLAN_VID_MASK|VLAN_CFI)
}

func (match *Match) SetVlanVidMasked(vid uint16, mask uint16) {
	mask &= VLAN_VID_MASK | VLAN_CFI
	ovid := match.Flow.VlanTci()
	match.Flow.SetVlanTci((ovid & ^mask) | (vid & mask))
	omask := match.Mask.VlanTci()
	match.Mask.SetVlanTci(mask | (omask & VLAN_PCP_MASK))
}

func (match *Match) SetDlType(t uint16) {
	match.Mask.SetDlType(math.MaxUint16)
	match.Flow.SetDlType(t)
}

func (match *Match) SetTunId(t uint64) {
	match.SetTunIdMasked(t, math.MaxUint64)
}

func (match *Match) SetTunIdMasked(t uint64, m uint64) {
	match.Mask.SetTunId(m)
	match.Flow.SetTunId(t)
}

func (match *Match) SetEth(p []byte, addr []byte, maskSet []byte) {
	copy(addr, p)
	for i := range maskSet {
		maskSet[i] = 0xff
	}
}

func (match *Match) SetEthMasked(p []byte, m []byte, set []byte, maskSet []byte) {
	for i := range maskSet {
		set[i] = p[i] & m[i]
		maskSet[i] = m[i]
	}
}

func (match *Match) SetMetadata(t uint64) {
	match.SetMetadataMasked(t, math.MaxUint64)
}

func (match *Match) SetMetadataMasked(t uint64, m uint64) {
	match.Mask.SetMetadata(m)
	match.Flow.SetMetadata(t)
}

func (match *Match) SetNwDscp(t uint8) {
	match.Mask.SetNwTos(match.Mask.NwTos() | IP_DSCP_MASK)
	tos := match.Flow.NwTos() & (0xff & ^IP_DSCP_MASK)
	match.Flow.SetNwTos(tos | (t & IP_DSCP_MASK))
}

func (match *Match) SetNwEcn(t uint8) {
}
