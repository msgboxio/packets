package packets

const (
	mpls_lse = 32
)

type mpls struct {
}

func (o mpls) Lse(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 0); err != nil {
		return
	}
	v &= ((1 << (mpls_lse + 0)) - 1) >> 0
	return
}
func (o mpls) SetLse(b []byte, vn uint32) error {
	v, err := read_b32(b, 0)
	if err != nil {
		return err
	}
	err = write_b32(b, 0, v|(vn<<0)&((1<<(mpls_lse+0))-1))
	return err
}
