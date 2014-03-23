package packets

const (
	nd_opt_ndtype = 8
	nd_opt_length = 8
)

type nd_opt struct {
}

func (o nd_opt) Ndtype(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (nd_opt_ndtype + 0)) - 1) >> 0
	return
}
func (o nd_opt) SetNdtype(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(nd_opt_ndtype+0))-1))
	return err
}
func (o nd_opt) Length(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (nd_opt_length + 0)) - 1) >> 0
	return
}
func (o nd_opt) SetLength(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(nd_opt_length+0))-1))
	return err
}
