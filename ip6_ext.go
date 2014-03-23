package packets

const (
	ip6_ext_next = 8
	ip6_ext_len  = 8
)

type ip6_ext struct {
}

func (o ip6_ext) Next(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (ip6_ext_next + 0)) - 1) >> 0
	return
}
func (o ip6_ext) SetNext(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(ip6_ext_next+0))-1))
	return err
}
func (o ip6_ext) Len(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (ip6_ext_len + 0)) - 1) >> 0
	return
}
func (o ip6_ext) SetLen(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(ip6_ext_len+0))-1))
	return err
}
