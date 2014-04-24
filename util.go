package packets

import (
	"bytes"
	"encoding/binary"
	"io"
)

func ReadB64(b []byte, offset int) (uint64, error) {
	if len(b) < 8+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint64(b[offset : offset+8]), nil
}
func ReadB32(b []byte, offset int) (uint32, error) {
	if len(b) < 4+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint32(b[offset : offset+4]), nil
}
func ReadB16(b []byte, offset int) (uint16, error) {
	if len(b) < 2+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint16(b[offset : offset+2]), nil
}
func ReadB8(b []byte, offset int) (uint8, error) {
	if len(b) < 1+offset {
		return 0, io.EOF
	}
	return b[offset], nil
}

func ReadSB32(b []byte, offset int) (int32, error) {
	if len(b) < 4+offset {
		return 0, io.EOF
	}
	return int32(binary.BigEndian.Uint32(b[offset : offset+4])), nil
}

// readSlice is like ReadBytes but returns a reference to internal buffer data.
func ReadSlice(b []byte, off int, lim int, delim byte) (line []byte, err error) {
	i := bytes.IndexByte(b[off:], delim)
	end := off + i
	if i < 0 {
		end = len(b)
		err = io.EOF
	}
	line = b[off:end]
	return line, err
}

func WriteB64(b []byte, offset int, v uint64) error {
	if len(b) < 8+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint64(b[offset:offset+8], v)
	return nil
}
func WriteB32(b []byte, offset int, v uint32) error {
	if len(b) < 4+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint32(b[offset:offset+4], v)
	return nil
}
func WriteB16(b []byte, offset int, v uint16) error {
	if len(b) < 2+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint16(b[offset:offset+2], v)
	return nil
}
func WriteB8(b []byte, offset int, v uint8) error {
	if len(b) < offset {
		return io.EOF
	}
	b[offset] = v
	return nil
}

func WriteSB32(b []byte, offset int, v int32) error {
	if len(b) < 4+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint32(b[offset:offset+4], uint32(v))
	return nil
}

//
// misc
//
var (
	// Integer.numberOfTrailingZeros
	// http://stackoverflow.com/questions/5471129/number-of-trailing-zeros
	// http://graphics.stanford.edu/~seander/bithacks.html#ZerosOnRightModLookup
	ZerosOnRightModLookup = []int{
		32, 0, 1, 26, 2, 23, 27, 0, 3, 16, 24, 30, 28, 11, 0, 13, 4, 7, 17,
		0, 25, 22, 31, 15, 29, 10, 12, 6, 0, 21, 14, 9, 5, 20, 8, 19, 18,
	}
)

// counts number of trailing zeros
func Ntz(i uint32) int {
	return ZerosOnRightModLookup[(i&-i)%37]
}

// eth
func EthAddrZero(ea []byte) bool {
	return (ea[0] | ea[1] | ea[2] | ea[3] | ea[4] | ea[5]) == 0
}
func EthMaskIsExact(ea []byte) bool {
	return (ea[0] & ea[1] & ea[2] & ea[3] & ea[4] & ea[5]) == 0xff
}

// ip4
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
var (
	in6addrAny   = [...]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
