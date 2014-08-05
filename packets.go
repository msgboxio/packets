package packets

import (
	"bytes"
	"net"
)

/* 802.1Q. */
const (
	QLAN_QTAG_LEN     = 4
	VLAN_VID_MASK     = uint16(0x0fff)
	VLAN_VID_SHIFT    = 0
	VLAN_PCP_MASK     = 0xe000
	VLAN_PCP_SHIFT    = 13
	VLAN_CFI          = 0x1000
	VLAN_CFI_SHIFT    = 12
	VLAN_VID_CFI_MASK = VLAN_CFI | VLAN_VID_MASK
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

// eth
func EthAddrZero(ea []byte) bool {
	return (ea[0] | ea[1] | ea[2] | ea[3] | ea[4] | ea[5]) == 0
}
func EthMaskIsExact(ea []byte) bool {
	return (ea[0] & ea[1] & ea[2] & ea[3] & ea[4] & ea[5]) == 0xff
}

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

/* Given a mpls label stack entry in network byte order
 * return mpls label in host byte order */
func MplsLseToLabel(lse uint32) uint32 {
	return (lse & MPLS_LABEL_MASK) >> MPLS_LABEL_SHIFT
}

/* Given a mpls label stack entry in network byte order
 * return mpls tc */
func MplsLseToTc(lse uint32) uint8 {
	return uint8((lse & MPLS_TC_MASK) >> MPLS_TC_SHIFT)
}

/* Given a mpls label stack entry in network byte order
 * return mpls BoS bit  */
func MplsLseToBos(lse uint32) uint8 {
	if lse&MPLS_BOS_MASK != 0 {
		return 1
	}
	return 0
}

// ip4
const (
	// TOS fields
	IP_ECN_NOT_ECT = uint8(0x0)
	IP_ECN_ECT_1   = 0x01
	IP_ECN_ECT_0   = 0x02
	IP_ECN_CE      = 0x03
	IP_ECN_MASK    = 0x03
	IP_DSCP_MASK   = 0xfc
	// Frag Fields
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

/* Returns true if 'netmask' is a CIDR netmask, that is, if it consists of N
 * high-order 1-bits and 32-N low-order 0-bits. */
func IsCidr(ip uint32) bool {
	x := 0xffffffff & ^ip
	return (x & (x + 1)) == 0
}

/* Given the IP netmask 'netmask', returns the number of bits of the IP address
 * that it specifies, that is, the number of 1-bits in 'netmask'.
 *
 * If 'netmask' is not a CIDR netmask (see ip_is_cidr()), the return value will
 * still be in the valid range but isn't otherwise meaningful. */
func CountCidrBits(netmask uint32) int {
	return 32 - Ntz(netmask)
}

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

// ip6
var (
	in6addrAny   = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	in6addrExact = []byte{0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff}
)

/* Returns true if 'netmask' is a CIDR netmask, that is, if it consists of N
 * high-order 1-bits and 128-N low-order 0-bits. */
func IsV6Cidr(netmask []byte) bool {
	for i := range netmask {
		if netmask[i] != 0xff {
			x := ^netmask[i]
			if (x & (x + 1)) != 0 {
				return false
			}
			for i < 16 {
				if netmask[i] != 0 {
					return false
				}
				i++
			}
		}
	}
	return true
}

func IsV6MaskAny(mask []byte) bool {
	return bytes.Equal(mask, in6addrAny[:])
}

func IsV6MaskExact(mask []byte) bool {
	return bytes.Equal(mask, in6addrExact[:])
}

/* Given the IPv6 netmask 'netmask', returns the number of bits of the IPv6
 * address that it specifies, that is, the number of 1-bits in 'netmask'.
 * 'netmask' must be a CIDR netmask (see ipv6_is_cidr()).
 *
 * If 'netmask' is not a CIDR netmask (see ipv6_is_cidr()), the return value
 * will still be in the valid range but isn't otherwise meaningful. */
func CountV6CidrBits(mask []byte) (count int) {
	for i := range mask {
		if mask[i] == 0xff {
			count += 8
		} else {
			for nm := mask[i]; nm != 0; nm <<= 1 {
				count++
			}
			break
		}
	}
	return
}

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

// hash
func hashRot(x, k uint32) uint32 {
	return (x << k) | (x >> (32 - k))
}

/* Murmurhash by Austin Appleby,
 * from http://code.google.com/p/smhasher/source/browse/trunk/MurmurHash3.cpp.
 *
 * The upstream license there says:
 *
 * // MurmurHash3 was written by Austin Appleby, and is placed in the public
 * // domain. The author hereby disclaims copyright to this source code.
 *
 * See hash_words() for sample usage. */

func mhashAdd_(hash, data uint32) uint32 {
	data *= 0xcc9e2d51
	data = hashRot(data, 15)
	data *= 0x1b873593
	return hash ^ data
}

func mhashAdd(hash, data uint32) uint32 {
	hash = mhashAdd_(hash, data)
	hash = hashRot(hash, 13)
	return hash*5 + 0xe6546b64
}

func mhashFinish(hash, nBytes uint32) uint32 {
	hash ^= nBytes
	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16
	return hash
}

/* Returns the hash of 'a', 'b', and 'c'. */
func hash3Words(a, b, c uint32) uint32 {
	return mhashFinish(mhashAdd(mhashAdd(mhashAdd(a, 0), b), c), 12)
}

func hashUint64Basis(x uint64, basis uint32) uint32 {
	return hash3Words(uint32(x>>32), uint32(x), basis)
}

func EthAddrToUint64(ea net.HardwareAddr) uint64 {
	return (uint64(ea[0]) << 40) |
		(uint64(ea[1]) << 32) |
		(uint64(ea[2]) << 24) |
		(uint64(ea[3]) << 16) |
		(uint64(ea[4]) << 8) |
		uint64(ea[5])
}

func EthAddrVlanToUint64(ea net.HardwareAddr, vlan uint16) uint64 {
	return ((uint64(vlan) << 48) | EthAddrToUint64(ea))
}

func HashMac(ea net.HardwareAddr, vlan uint16, basis uint32) uint32 {
	return hashUint64Basis(EthAddrVlanToUint64(ea, vlan), basis)
}
