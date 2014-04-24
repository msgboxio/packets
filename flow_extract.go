package packets

import (
	"bytes"
)

const (
	FLOW_DL_TYPE_NONE = 0x5ff

	/* Fragment bits, used for IPv4 and IPv6, always zero for non-IP flows. */
	FLOW_NW_FRAG_ANY   = uint8(1 << 0) /* Set for any IP frag. */
	FLOW_NW_FRAG_LATER = (1 << 1)      /* Set for IP frag with nonzero offset. */
	FLOW_NW_FRAG_MASK  = (FLOW_NW_FRAG_ANY | FLOW_NW_FRAG_LATER)

	FLOW_TNL_F_DONT_FRAGMENT = (1 << 0)
	FLOW_TNL_F_CSUM          = (1 << 1)
	FLOW_TNL_F_KEY           = (1 << 2)
)

func (flow *Flow) Extract(b []byte) {
	if len(b) < ETH_HEADER_LEN {
		return
	}
	buf := bytes.NewBuffer(b)
	buf.Read(flow.DlDst())
	buf.Read(flow.DlSrc())
	// l2
	ty, _ := ReadB16(buf.Bytes(), 0) //peek
	if ty == ETH_TYPE_VLAN {
		b := buf.Bytes()
		if len(b) < QLAN_QTAG_LEN {
			return
		}
		var qtag VlanQtag
		flow.SetVlanTci(qtag.Tci(b) | VLAN_CFI)
	} else {
		ty = extractLlcSnap(buf)
	}
	flow.SetDlType(ty)
	switch ty {
	case ETH_TYPE_MPLS:
		flow.ExtractMpls(buf) //2.5
	case ETH_TYPE_IP:
		flow.ExtractIp(buf) //l3
	case ETH_TYPE_IPV6:
		flow.ExtractIp6(buf) //l3
	case ETH_TYPE_ARP, ETH_TYPE_RARP:
		flow.ExtractArp(buf) //l3
	}
}

// TODO - single MPLS header supported only
func (flow *Flow) ExtractMpls(buf *bytes.Buffer) {
	var mpls Mpls
	for {
		b := buf.Bytes()
		if len(b) < MPLS_HEADER_LEN {
			break
		}
		lse := mpls.Lse(b)
		if (lse & MPLS_BOS_MASK) != 0 {
			break
		}
		flow.SetMplsLse(lse)
		buf.Next(MPLS_HEADER_LEN)
	}
}

func (flow *Flow) ExtractIp(buf *bytes.Buffer) {
	if buf.Len() < IP_HEADER_LEN {
		return
	}
	var ip Ip4
	b := buf.Bytes()
	ipLen := ip.Hdrlen(b) * 4
	if (ipLen >= IP_HEADER_LEN) && (len(b) >= int(ipLen)) {
		buf.Next(int(ipLen)) // advance buffer
		// ip4 header
		protocol := ip.Protocol(b)

		flow.SetNwTos(ip.Tos(b))
		flow.SetNwProto(protocol)
		flow.SetNwSrc(ip.Srcaddr(b))
		flow.SetNwDst(ip.Dstaddr(b))
		flow.SetNwTtl(ip.Ttl(b))

		offset := ip.Offset(b)
		if isFrag(offset) {
			f := FLOW_NW_FRAG_ANY
			if (offset & IP_FRAG_OFF_MASK) != 0 {
				f |= FLOW_NW_FRAG_LATER
			}
			flow.SetNwFrag(f)
		}
		if (offset & IP_FRAG_OFF_MASK) == 0 {
			switch protocol {
			case IPPROTO_TCP:
				flow.ExtractTcp(buf)
			case IPPROTO_UDP:
				flow.ExtractUdp(buf)
			case IPPROTO_SCTP:
				flow.ExtractSctp(buf)
			case IPPROTO_ICMP:
				flow.ExtractIcmp(buf)
			}
		}
	}
}

func (flow *Flow) ExtractIp6(buf *bytes.Buffer) {
	if buf.Len() < IP6_HEADER_LEN {
		return
	}
	b := buf.Next(IP6_HEADER_LEN) // extract header
	var ip Ip6
	flow.SetIpv6Src(ip.Source(b))
	flow.SetIpv6Dst(ip.Dest(b))
	v6Flow := ip.Flow(b)
	flow.SetNwTos(uint8(v6Flow >> 20))
	flow.SetIpv6Label(v6Flow & IPV6_LABEL_MASK)
	hl := ip.HopLimit(b)
	flow.SetNwTtl(hl)
	flow.SetNwProto(IPPROTO_NONE)
	nexthdr := ip.Next(b)
	var hdrlen int
	var ip6_ext Ip6Ext
	for {
		if (nexthdr != IPPROTO_HOPOPTS) &&
			(nexthdr != IPPROTO_ROUTING) &&
			(nexthdr != IPPROTO_DSTOPTS) &&
			(nexthdr != IPPROTO_AH) &&
			(nexthdr != IPPROTO_FRAGMENT) {
			/* It's either a terminal header (e.g., TCP, UDP) or one we
			 * don't understand.  In either case, we're done with the
			 * packet, so use it to fill in 'nw_proto'. */
			break
		}
		if buf.Len() < IP6_EXT_HEADER_MIN_LEN {
			return // EINVAL
		}

		if (nexthdr == IPPROTO_HOPOPTS) ||
			(nexthdr == IPPROTO_ROUTING) ||
			(nexthdr == IPPROTO_DSTOPTS) {
			/* These headers, while different, have the fields we care about
			 * in the same location and with the same interpretation. */
			extb := buf.Next(IP6_EXT_HEADER_LEN)
			nexthdr = ip6_ext.Next(extb)
			hdrlen_ := ip6_ext.Len(extb)
			hdrlen = 8 * (int(hdrlen_) + 1)
			if buf.Len() < hdrlen {
				return // EINVAL;
			}
			buf.Next(hdrlen - IP6_EXT_HEADER_LEN) // extract nxtHeader
		} else if nexthdr == IPPROTO_AH {
			/* A standard AH definition isn't available, but the fields
			 * we care about are in the same location as the generic
			 * option header--only the header length is calculated
			 * differently. */
			extb := buf.Next(IP6_EXT_HEADER_LEN)
			nexthdr = ip6_ext.Next(extb)
			hdrlen_ := ip6_ext.Len(extb)
			hdrlen = 4 * (int(hdrlen_) + 2)
			if buf.Len() < hdrlen {
				return // EINVAL;
			}
			buf.Next(hdrlen - IP6_EXT_HEADER_LEN) // extract nxtHeader
		} else if nexthdr == IPPROTO_FRAGMENT {
			// IP6_EXT_FRAG_HEADER_LEN is 8, so already checked min len
			extb := buf.Next(IP6_EXT_FRAG_HEADER_LEN)
			nexthdr = ip6_ext.Next(extb)
			/* We only process the first fragment. */
			var ip6Frag Ip6FragHeader
			fragment := ip6Frag.Fragment(extb)
			if fragment != 0 {
				f := FLOW_NW_FRAG_ANY
				if (fragment & IP6F_OFF_MASK) != 0 {
					f |= FLOW_NW_FRAG_LATER
					nexthdr = IPPROTO_FRAGMENT
				}
				flow.SetNwFrag(f)
				break
			}
		} // while loop
		flow.SetNwProto(nexthdr)
	}
}

func (flow *Flow) ExtractArp(buf *bytes.Buffer) {
	if buf.Len() < ARP_ETH_HEADER_LEN {
		return
	}
	b := buf.Next(ARP_ETH_HEADER_LEN) // extract header
	var arp Arp
	if (arp.Hrd(b) == 1) &&
		(arp.Pro(b) == ETH_TYPE_IP) &&
		(arp.Hln(b) == ETH_ADDR_LEN) &&
		(arp.Pln(b) == 4) {
		/* We only match on the lower 8 bits of the opcode. */
		if op := arp.Op(b); op <= 0xff {
			flow.SetNwProto(uint8(op))
		}
		flow.SetNwSrc(arp.Spa(b))
		flow.SetNwDst(arp.Tpa(b))
		flow.SetArpSha(arp.Sha(b))
		flow.SetArpTha(arp.Tha(b))
	}
}

func (flow *Flow) ExtractTcp(buf *bytes.Buffer) {
	if buf.Len() < TCP_HEADER_LEN {
		return
	}
	b := buf.Next(TCP_HEADER_LEN) // extract header
	var tcp Tcp
	ctl := tcp.Ctl(b)
	tcpLen := tcpOffset(ctl)
	if tcpLen < TCP_HEADER_LEN && buf.Len() < int(tcpLen) {
		return
	}
	tcpb := buf.Next(int(tcpLen - TCP_HEADER_LEN))
	flow.SetTpSrc(tcp.Src(tcpb))
	flow.SetTpDst(tcp.Dst(tcpb))
	flow.SetTcpFlags(ctl & 0xff)
}

func (flow *Flow) ExtractUdp(buf *bytes.Buffer) {
	if buf.Len() < UDP_HEADER_LEN {
		return
	}
	b := buf.Next(UDP_HEADER_LEN) // extract header
	var udp Udp
	flow.SetTpSrc(udp.Src(b))
	flow.SetTpDst(udp.Dst(b))
}

func (flow *Flow) ExtractSctp(buf *bytes.Buffer) {
	if buf.Len() < SCTP_HEADER_LEN {
		return
	}
	b := buf.Next(SCTP_HEADER_LEN) // extract header
	var sctp Sctp
	flow.SetTpSrc(sctp.Src(b))
	flow.SetTpDst(sctp.Dst(b))
}

func (flow *Flow) ExtractIcmp(buf *bytes.Buffer) {
	if buf.Len() < ICMP_HEADER_LEN {
		return
	}
	b := buf.Next(ICMP_HEADER_LEN) // extract header
	var icmp Icmp4
	/* ICMPv6 type and code fields use the 16-bit transport port fields*/
	flow.SetTpSrc(uint16(icmp.Type(b)))
	flow.SetTpDst(uint16(icmp.Code(b)))
}

func (flow *Flow) ExtractIcmp6(buf *bytes.Buffer) {
	if buf.Len() < ICMPV6_HEADER_LEN {
		return
	}
	b := buf.Next(ICMPV6_HEADER_LEN) // extract header
	var icmp Icmp6
	/* ICMPv6 type and code fields use the 16-bit transport port fields*/
	src := icmp.Type(b)
	dst := icmp.Code(b)
	flow.SetTpSrc(uint16(src))
	flow.SetTpDst(uint16(dst))
	// solicitation
	if (dst == 0) &&
		((src == ND_NEIGHBOR_SOLICIT) || (src == ND_NEIGHBOR_ADVERT)) {
		if buf.Len() < IP6_ADDR_LEN {
			return
		}
		flow.SetNdTarget(buf.Next(IP6_ADDR_LEN))
		for buf.Len() >= NDOPT_HEADER_LEN {
			/* The minimum size of an option is 8 bytes, which also is
			 * the size of Ethernet link-layer options. */
			ndOptB := buf.Next(NDOPT_HEADER_LEN)
			var opt NdOpt
			optLen := (opt.Length(ndOptB) * 8) - 2
			if int(optLen) < buf.Len() {
				break
			}
			/* Store the link layer address if the appropriate option is
			 * provided.  It is considered an error if the same link
			 * layer option is specified twice. */
			optType := opt.Ndtype(ndOptB)
			if (optType == ND_OPT_SOURCE_LINKADDR) && (optLen == 8) {
				if EthAddrZero(flow.ArpSha()) {
					flow.SetArpSha(buf.Next(ETH_ADDR_LEN))
					optLen -= ETH_ADDR_LEN
				} else {
					break
				}
			} else if (optType == ND_OPT_TARGET_LINKADDR) && (optLen == 8) {
				if EthAddrZero(flow.ArpTha()) {
					flow.SetArpTha(buf.Next(ETH_ADDR_LEN))
					optLen -= ETH_ADDR_LEN
				} else {
					break
				}
			}
			buf.Next(int(optLen))
		} // for
	} // solicit
}

func (flow *Flow) ExtractX(buf *bytes.Buffer) {}

func isFrag(ip_frag_off uint16) bool {
	return (ip_frag_off & (IP_MORE_FRAGMENTS | IP_FRAG_OFF_MASK)) != 0
}

func tcpOffset(ctl uint16) uint16 {
	return ctl >> 12
}

func extractLlcSnap(buf *bytes.Buffer) uint16 {
	ty, _ := ReadB16(buf.Next(2), 0)
	if ty >= ETH_TYPE_MIN {
		return ty
	}
	b := buf.Bytes()
	if len(b) < LLC_SNAP_HEADER_LEN {
		return FLOW_DL_TYPE_NONE
	}

	var llc Llc
	dsap := llc.Dsap(b)
	ssap := llc.Ssap(b)
	cntl := llc.Cntl(b)
	if (dsap != LLC_DSAP_SNAP) && (ssap != LLC_SSAP_SNAP) && (cntl != LLC_CNTL_SNAP) {
		return FLOW_DL_TYPE_NONE
	}
	buf.Next(LLC_HEADER_LEN)

	var snap Snap
	b = buf.Bytes()
	org := snap.Org(b)
	snapt := snap.Type(b)
	if (org != 0) && (ty >= ETH_TYPE_MIN) {
		// header is present
		buf.Next(SNAP_HEADER_LEN)
		return snapt
	}
	return FLOW_DL_TYPE_NONE
}
