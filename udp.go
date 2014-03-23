package packets

const (
	udp_src  = 16
	udp_dst  = 16
	udp_len  = 16
	udp_csum = 16
)

type udp struct {
}

func (o udp) Src(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 0); err != nil {
		return
	}
	v &= ((1 << (udp_src + 0)) - 1) >> 0
	return
}
func (o udp) SetSrc(b []byte, vn uint16) error {
	v, err := read_b16(b, 0)
	if err != nil {
		return err
	}
	err = write_b16(b, 0, v|(vn<<0)&((1<<(udp_src+0))-1))
	return err
}
func (o udp) Dst(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (udp_dst + 0)) - 1) >> 0
	return
}
func (o udp) SetDst(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(udp_dst+0))-1))
	return err
}
func (o udp) Len(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 4); err != nil {
		return
	}
	v &= ((1 << (udp_len + 0)) - 1) >> 0
	return
}
func (o udp) SetLen(b []byte, vn uint16) error {
	v, err := read_b16(b, 4)
	if err != nil {
		return err
	}
	err = write_b16(b, 4, v|(vn<<0)&((1<<(udp_len+0))-1))
	return err
}
func (o udp) Csum(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 6); err != nil {
		return
	}
	v &= ((1 << (udp_csum + 0)) - 1) >> 0
	return
}
func (o udp) SetCsum(b []byte, vn uint16) error {
	v, err := read_b16(b, 6)
	if err != nil {
		return err
	}
	err = write_b16(b, 6, v|(vn<<0)&((1<<(udp_csum+0))-1))
	return err
}
