package packets

import (
	"fmt"
	"math"
	"net"
	"strings"
)

type Match struct {
	Flow *Flow
	Mask *Flow
}

func NewMatch() *Match {
	return &Match{NewFlow(), NewFlow()}
}

func NewMatchFromFlow(flow *Flow) *Match {
	match := &Match{flow, NewFlow()}
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
	mask := match.Mask
	mask.SetDlType(0xffff)
	if flow.NwProto() != 0 {
		mask.SetNwProto(0xff)
	}
	if flow.SkbPriority() != 0 {
		mask.SetSkbPriority(math.MaxUint32)
	}
	if flow.PktMark() != 0 {
		mask.SetPktMark(math.MaxUint32)
	}
	// for i := 0; i < FLOW_N_REGS; i++ {
	// 	if flow.u32_regs[i] != 0 {
	// 		mask.regs[1] = math.MaxUint32
	// 	}
	// }
	if flow.TunIpDst() != 0 {
		if (flow.TunFlags() & FLOW_TNL_F_KEY) != 0 {
			mask.SetTunId(math.MaxUint64)
		}
		mask.SetTunIpSrc(math.MaxUint32)
		mask.SetTunIpDst(math.MaxUint32)
		mask.SetTunFlags(math.MaxUint16)
		mask.SetTunIpTos(math.MaxUint8)
		mask.SetTunIpTtl(math.MaxUint8)
	} else if flow.TunId() != 0 {
		mask.SetTunId(math.MaxUint64)
	}
	mask.SetMetadata(math.MaxUint64)
	mask.SetInPort(math.MaxUint16)
	mask.SetVlanTci(math.MaxUint16)
	fill(mask.DlSrc(), math.MaxUint8)
	fill(mask.DlDst(), math.MaxUint8)

	dl := flow.DlType()
	if dl == ETH_TYPE_IPV6 {
		fill(mask.Ipv6Src(), math.MaxUint8)
		fill(mask.Ipv6Dst(), math.MaxUint8)
		mask.SetIpv6Label(math.MaxUint32)
	} else if dl == ETH_TYPE_IP || dl == ETH_TYPE_ARP || dl == ETH_TYPE_RARP {
		mask.SetNwSrc(math.MaxUint32)
		mask.SetNwDst(math.MaxUint32)
	} else if dl == ETH_TYPE_MPLS || dl == ETH_TYPE_MPLS_MCAST {
		mask.SetMplsLse(math.MaxUint32)
	}
	if dl == ETH_TYPE_ARP || dl == ETH_TYPE_RARP {
		fill(mask.ArpSha(), math.MaxUint8)
		fill(mask.ArpTha(), math.MaxUint8)
	}
	if dl == ETH_TYPE_IP || dl == ETH_TYPE_IPV6 {
		mask.SetNwTos(math.MaxUint8)
		mask.SetNwTtl(math.MaxUint8)
		if flow.NwFrag() != 0 {
			mask.SetNwFrag(math.MaxUint8)
			if (flow.NwFrag() & FLOW_NW_FRAG_LATER) != 0 {
				/* No transport layer header in later fragments. */
				return
			}
		}
		nw := flow.NwProto()
		if nw == IPPROTO_ICMP || nw == IPPROTO_ICMPV6 ||
			(flow.TpSrc() != 0 || flow.TpDst() != 0) {
			mask.SetTpSrc(math.MaxUint16)
			mask.SetTpDst(math.MaxUint16)
		}
		if nw == IPPROTO_TCP {
			mask.SetTcpFlags(math.MaxUint16)
		}
		if nw == IPPROTO_ICMPV6 {
			fill(mask.ArpSha(), math.MaxUint8)
			fill(mask.ArpTha(), math.MaxUint8)
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

func (match *Match) SetNwDscp(dscp uint8) {
	match.Mask.SetNwTos(match.Mask.NwTos() | IP_DSCP_MASK)
	tos := match.Flow.NwTos() & (0xff & ^IP_DSCP_MASK)
	match.Flow.SetNwTos(tos | (dscp & IP_DSCP_MASK))
}

func (match *Match) SetNwEcn(nw_ecn uint8) {
	match.Mask.SetNwTos(match.Mask.NwTos() | IP_ECN_MASK)
	tos := match.Flow.NwTos() & (0xff & ^IP_ECN_MASK)
	match.Flow.SetNwTos(tos | (nw_ecn & IP_ECN_MASK))
}

func (match *Match) SetNwProto(nw_proto uint8) {
	match.Flow.SetNwProto(nw_proto)
	match.Mask.SetNwProto(math.MaxUint8)
}

func (match *Match) SetNwTtl(nw_ttl uint8) {
	match.Mask.SetNwTtl(math.MaxUint8)
	match.Flow.SetNwTtl(nw_ttl)
}

func (match *Match) SetNwFrag(nw_frag uint8) {
	match.Mask.SetNwFrag(match.Mask.NwFrag() | FLOW_NW_FRAG_MASK)
	match.Flow.SetNwFrag(nw_frag)
}

func (match *Match) SetIcmpType(icmp_type uint8) {
	match.Flow.SetTpSrc(uint16(icmp_type) & math.MaxUint16)
	match.Mask.SetTpSrc(math.MaxUint16)
}

func (f *Flow) String() string {
	match := NewMatchFromFlow(f)
	return match.String()
}

func (match *Match) String() (s string) {
	dltype := match.Flow.DlType()
	if match.Mask.DlType() != 0 {
		switch dltype {
		case ETH_TYPE_IP:
			if match.Mask.NwProto() != 0 {
				switch match.Flow.NwProto() {
				case IPPROTO_ICMP:
					s += "icmp,"
				case IPPROTO_TCP:
					s += "tcp,"
				case IPPROTO_UDP:
					s += "udp,"
				case IPPROTO_SCTP:
					s += "sctp,"
				default:
					s += "ip,"
				}
			} else {
				s += "ip,"
			}
		case ETH_TYPE_IPV6:
			if match.Mask.NwProto() != 0 {
				switch match.Flow.NwProto() {
				case IPPROTO_ICMPV6:
					s += "icmp6,"
				case IPPROTO_TCP:
					s += "tcp6,"
				case IPPROTO_UDP:
					s += "udp6,"
				case IPPROTO_SCTP:
					s += "sctp6,"
				default:
					s += "ipv6,"
				}
			} else {
				s += "ipv6,"
			}
		case ETH_TYPE_ARP:
			s += "arp,"
		case ETH_TYPE_RARP:
			s += "rarp,"
		case ETH_TYPE_MPLS:
			s += "mpls,"
		case ETH_TYPE_MPLS_MCAST:
			s += "mplsm,"
		}
	} // dltype
	s += formatMaskedU64("metadata", match.Flow.Metadata(), match.Mask.Metadata())
	if mask := match.Mask.InPort(); mask != 0 {
		s += fmt.Sprintf("in_port=%d,", match.Flow.InPort())
	}
	if mask := match.Mask.VlanTci(); mask != 0 {
		tci := match.Flow.VlanTci()
		if mask&VLAN_VID_MASK != 0 {
			s += fmt.Sprintf("vid=%d,", tci&VLAN_VID_MASK)
		}
	}
	s += formatMaskedEth("dl_src", match.Flow.DlSrc(), match.Mask.DlSrc())
	s += formatMaskedEth("dl_dst", match.Flow.DlDst(), match.Mask.DlDst())
	//
	switch dltype {
	case ETH_TYPE_IPV6:
		s += formatIp6Netmask("ipv6_src", match.Flow.Ipv6Src(), match.Mask.Ipv6Src())
		s += formatIp6Netmask("ipv6_dst", match.Flow.Ipv6Dst(), match.Mask.Ipv6Dst())
	case ETH_TYPE_ARP, ETH_TYPE_RARP:
		s += formatIpNetmask("arp_spa", match.Flow.NwSrc(), match.Mask.NwSrc())
		s += formatIpNetmask("arp_tpa", match.Flow.NwDst(), match.Mask.NwDst())
	default:
		s += formatIpNetmask("nw_src", match.Flow.NwSrc(), match.Mask.NwSrc())
		s += formatIpNetmask("nw_dst", match.Flow.NwDst(), match.Mask.NwDst())
	}
	//
	if match.Mask.NwProto() != 0 {
		switch dltype {
		case ETH_TYPE_ARP, ETH_TYPE_RARP:
			s += fmt.Sprintf("arp_op=%d,", match.Flow.NwProto())
		default:
			s += fmt.Sprintf("nw_proto=%d,", match.Flow.NwProto())
		}
	}
	if dltype == ETH_TYPE_ARP || dltype == ETH_TYPE_RARP {
		s += formatMaskedEth("arp_sha", match.Flow.ArpSha(), match.Mask.ArpSha())
		s += formatMaskedEth("arp_tha", match.Flow.ArpTha(), match.Mask.ArpTha())
	}

	if (match.Mask.NwTos() & IP_DSCP_MASK) != 0 {
		s += fmt.Sprintf("nw_tos=%d,", match.Flow.NwTos()&IP_DSCP_MASK)
	}
	if (match.Mask.NwTos() & IP_ECN_MASK) != 0 {
		s += fmt.Sprintf("nw_ecn=%d,", match.Flow.NwTos()&IP_ECN_MASK)
	}

	s = strings.TrimSuffix(s, ",")
	return
}

func formatIp(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		(ip>>24)&0xff,
		(ip>>16)&0xff,
		(ip>>8)&0xff,
		ip&0xff)
}

func formatIpNetmask(f string, ip uint32, mask uint32) (s string) {
	if mask == 0 {
		return
	}
	s = fmt.Sprintf("%s=%s", f, formatIp(ip))
	if mask != math.MaxUint32 {
		if IsCidr(mask) {
			s += fmt.Sprintf("/%s", CountCidrBits(mask))
		} else {
			s += fmt.Sprintf("/%s", formatIp(mask))
		}
	}
	s += ","
	return
}

func formatIp6(ip []byte) string {
	return net.IP(ip).String()
}

func formatIp6Netmask(f string, v []byte, mask []byte) (s string) {
	if !IsV6MaskAny(mask) {
		s = fmt.Sprintf("%s=%s,", f, formatIp6(v))
		if mask != nil && !IsV6MaskExact(mask) {
			if IsV6Cidr(mask) {
				s += fmt.Sprintf("/%s", CountV6CidrBits(mask))
			} else {
				s += fmt.Sprintf("/%s", formatIp6(mask))
			}
			s += ","
		}
	}
	return
}

func formatMaskedU64(f string, v uint64, m uint64) (s string) {
	if m != 0 {
		s = fmt.Sprintf("%s=%#x", f, v)
		if m != math.MaxUint64 {
			s += fmt.Sprintf("/%#x", m)
		}
		s += ","
	}
	return
}

func formatEthMask(v []byte, m []byte) (s string) {
	s = net.HardwareAddr(v).String()
	if !EthMaskIsExact(m) {
		s += "/"
		s += net.HardwareAddr(m).String()
	}
	return
}

func formatMaskedEth(f string, v []byte, m []byte) (s string) {
	if !EthAddrZero(m) {
		s = fmt.Sprintf("%s=%s,", f, formatEthMask(v, m))
	}
	return
}
