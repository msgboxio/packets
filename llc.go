package packets

const (
	llc_dsap = 8
	llc_ssap = 8
	llc_cntl = 8
)

type llc struct {
}

func (o llc) Dsap(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (llc_dsap + 0)) - 1) >> 0
	return
}
func (o llc) SetDsap(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(llc_dsap+0))-1))
	return err
}
func (o llc) Ssap(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (llc_ssap + 0)) - 1) >> 0
	return
}
func (o llc) SetSsap(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(llc_ssap+0))-1))
	return err
}
func (o llc) Cntl(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 2); err != nil {
		return
	}
	v &= ((1 << (llc_cntl + 0)) - 1) >> 0
	return
}
func (o llc) SetCntl(b []byte, vn uint8) error {
	v, err := read_b8(b, 2)
	if err != nil {
		return err
	}
	err = write_b8(b, 2, v|(vn<<0)&((1<<(llc_cntl+0))-1))
	return err
}
