package packets

const (
	ip4_version  = 4
	ip4_hdrLen   = 4
	ip4_tos      = 8
	ip4_length   = 16
	ip4_id       = 16
	ip4_flags    = 3
	ip4_offset   = 13
	ip4_ttl      = 8
	ip4_protocol = 8
	ip4_hdrChks  = 16
	ip4_srcAddr  = 32
	ip4_dstAddr  = 32
	ip4_options  = '*'
)

type ip4 struct {
}

func (o ip4) Version(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (ip4_version + 0)) - 1) >> 0
	return
}
func (o ip4) SetVersion(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(ip4_version+0))-1))
	return err
}
func (o ip4) HdrLen(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (ip4_hdrLen + 4)) - 1) >> 4
	return
}
func (o ip4) SetHdrLen(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<4)&((1<<(ip4_hdrLen+4))-1))
	return err
}
func (o ip4) Tos(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (ip4_tos + 0)) - 1) >> 0
	return
}
func (o ip4) SetTos(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(ip4_tos+0))-1))
	return err
}
func (o ip4) Length(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (ip4_length + 0)) - 1) >> 0
	return
}
func (o ip4) SetLength(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(ip4_length+0))-1))
	return err
}
func (o ip4) Id(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 4); err != nil {
		return
	}
	v &= ((1 << (ip4_id + 0)) - 1) >> 0
	return
}
func (o ip4) SetId(b []byte, vn uint16) error {
	v, err := read_b16(b, 4)
	if err != nil {
		return err
	}
	err = write_b16(b, 4, v|(vn<<0)&((1<<(ip4_id+0))-1))
	return err
}
func (o ip4) Flags(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 6); err != nil {
		return
	}
	v &= ((1 << (ip4_flags + 0)) - 1) >> 0
	return
}
func (o ip4) SetFlags(b []byte, vn uint8) error {
	v, err := read_b8(b, 6)
	if err != nil {
		return err
	}
	err = write_b8(b, 6, v|(vn<<0)&((1<<(ip4_flags+0))-1))
	return err
}
func (o ip4) Offset(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 6); err != nil {
		return
	}
	v &= ((1 << (ip4_offset + 3)) - 1) >> 3
	return
}
func (o ip4) SetOffset(b []byte, vn uint16) error {
	v, err := read_b16(b, 6)
	if err != nil {
		return err
	}
	err = write_b16(b, 6, v|(vn<<3)&((1<<(ip4_offset+3))-1))
	return err
}
func (o ip4) Ttl(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 8); err != nil {
		return
	}
	v &= ((1 << (ip4_ttl + 0)) - 1) >> 0
	return
}
func (o ip4) SetTtl(b []byte, vn uint8) error {
	v, err := read_b8(b, 8)
	if err != nil {
		return err
	}
	err = write_b8(b, 8, v|(vn<<0)&((1<<(ip4_ttl+0))-1))
	return err
}
func (o ip4) Protocol(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 9); err != nil {
		return
	}
	v &= ((1 << (ip4_protocol + 0)) - 1) >> 0
	return
}
func (o ip4) SetProtocol(b []byte, vn uint8) error {
	v, err := read_b8(b, 9)
	if err != nil {
		return err
	}
	err = write_b8(b, 9, v|(vn<<0)&((1<<(ip4_protocol+0))-1))
	return err
}
func (o ip4) HdrChks(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 10); err != nil {
		return
	}
	v &= ((1 << (ip4_hdrChks + 0)) - 1) >> 0
	return
}
func (o ip4) SetHdrChks(b []byte, vn uint16) error {
	v, err := read_b16(b, 10)
	if err != nil {
		return err
	}
	err = write_b16(b, 10, v|(vn<<0)&((1<<(ip4_hdrChks+0))-1))
	return err
}
func (o ip4) SrcAddr(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 12); err != nil {
		return
	}
	v &= ((1 << (ip4_srcAddr + 0)) - 1) >> 0
	return
}
func (o ip4) SetSrcAddr(b []byte, vn uint32) error {
	v, err := read_b32(b, 12)
	if err != nil {
		return err
	}
	err = write_b32(b, 12, v|(vn<<0)&((1<<(ip4_srcAddr+0))-1))
	return err
}
func (o ip4) DstAddr(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 16); err != nil {
		return
	}
	v &= ((1 << (ip4_dstAddr + 0)) - 1) >> 0
	return
}
func (o ip4) SetDstAddr(b []byte, vn uint32) error {
	v, err := read_b32(b, 16)
	if err != nil {
		return err
	}
	err = write_b32(b, 16, v|(vn<<0)&((1<<(ip4_dstAddr+0))-1))
	return err
}
