package packets

const (
	QLAN_QTAG_LEN  = 4
	VLAN_VID_MASK  = uint16(0x0fff)
	VLAN_VID_SHIFT = 0
	VLAN_PCP_MASK  = 0xe000
	VLAN_PCP_SHIFT = 13
	VLAN_CFI       = 0x1000
	VLAN_CFI_SHIFT = 12
)

const (
	/*
	 * Minimum value for an Ethernet type.
	 * Values below this are IEEE 802.2 frame lengths.
	 */
	ETH_TYPE_MIN         = uint16(0x600)
	ETH_TYPE_IP          = 0x0800
	ETH_TYPE_ARP         = 0x0806
	ETH_TYPE_VLAN_8021Q  = 0x8100
	ETH_TYPE_VLAN        = ETH_TYPE_VLAN_8021Q
	ETH_TYPE_VLAN_8021AD = 0x88a8
	ETH_TYPE_IPV6        = 0x86dd
	ETH_TYPE_LACP        = 0x8809
	ETH_TYPE_RARP        = 0x8035
	ETH_TYPE_MPLS        = 0x8847
	ETH_TYPE_MPLS_MCAST  = 0x8848
)

const (
	ETH_ADDR_LEN    = 6
	ETH_HEADER_LEN  = 14
	ETH_PAYLOAD_MIN = 46
	ETH_PAYLOAD_MAX = 1500
)

// llc, snap
const (
	LLC_DSAP_SNAP = 0xaa
	LLC_SSAP_SNAP = 0xaa
	LLC_CNTL_SNAP = 3

	LLC_HEADER_LEN      = 3
	SNAP_HEADER_LEN     = 5
	LLC_SNAP_HEADER_LEN = (LLC_HEADER_LEN + SNAP_HEADER_LEN)
)

// mpls
const (
	MPLS_HEADER_LEN  = 4
	MPLS_TTL_MASK    = 0x000000ff
	MPLS_TTL_SHIFT   = 0
	MPLS_BOS_MASK    = 0x00000100
	MPLS_BOS_SHIFT   = 8
	MPLS_TC_MASK     = 0x00000e00
	MPLS_TC_SHIFT    = 9
	MPLS_LABEL_MASK  = 0xfffff000
	MPLS_LABEL_SHIFT = 12
)

// TOS fields
const (
	IP_ECN_NOT_ECT = uint8(0x0)
	IP_ECN_ECT_1   = 0x01
	IP_ECN_ECT_0   = 0x02
	IP_ECN_CE      = 0x03
	IP_ECN_MASK    = 0x03
	IP_DSCP_MASK   = 0xfc
)

// Frag Fields
const (
	IP_DONT_FRAGMENT  = uint16(0x4000) /* Don't fragment. */
	IP_MORE_FRAGMENTS = 0x2000         /* More fragments. */
	IP_FRAG_OFF_MASK  = 0x1fff         /* Fragment offset. */
)

const (
	ARP_ETH_HEADER_LEN = 28
)

const (
	IP_HEADER_LEN     = 20
	ICMP_HEADER_LEN   = 8
	ICMPV6_HEADER_LEN = 8
	SCTP_HEADER_LEN   = 12
	UDP_HEADER_LEN    = 8
)

const (
	IPPROTO_IP       = 0
	IPPROTO_HOPOPTS  = 0
	IPPROTO_ICMP     = 1
	IPPROTO_TCP      = 6
	IPPROTO_UDP      = 17
	IPPROTO_ROUTING  = 43
	IPPROTO_FRAGMENT = 44
	IPPROTO_AH       = 51
	IPPROTO_ICMPV6   = 58
	IPPROTO_NONE     = 59
	IPPROTO_DSTOPTS  = 60
	IPPROTO_SCTP     = 132
)

// ip6
const (
	IP6_HEADER_LEN          = 40
	IP6_ADDR_LEN            = 16
	IP6_EXT_HEADER_LEN      = 2
	IP6_EXT_HEADER_MIN_LEN  = 8
	IP6_EXT_FRAG_HEADER_LEN = 8

	NDOPT_HEADER_LEN = 2

	IP6F_OFF_MASK   = 0xfff8 /* mask out offset from _offlg */
	IPV6_LABEL_MASK = 0x000fffff

	ND_NEIGHBOR_SOLICIT = 135
	ND_NEIGHBOR_ADVERT  = 136

	ND_OPT_SOURCE_LINKADDR    = 1
	ND_OPT_TARGET_LINKADDR    = 2
	ND_OPT_PREFIX_INFORMATION = 3
	ND_OPT_REDIRECTED_HEADER  = 4
	ND_OPT_MTU                = 5
)

const (
	TCP_HEADER_LEN = 20
	TCP_FIN        = 0x001
	TCP_SYN        = 0x002
	TCP_RST        = 0x004
	TCP_PSH        = 0x008
	TCP_ACK        = 0x010
	TCP_URG        = 0x020
	TCP_ECE        = 0x040
	TCP_CWR        = 0x080
	TCP_NS         = 0x100
)
