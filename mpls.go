package packets

const (
	MplsLse = 32
)

type Mpls struct {
}

func (o Mpls) Lse(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 0); err != nil {
		return
	}
	v &= ((1 << (MplsLse + 0)) - 1) >> 0
	return
}
func (o Mpls) SetLse(b []byte, vn uint32) error {
	v, err := ReadB32(b, 0)
	if err != nil {
		return err
	}
	err = WriteB32(b, 0, v|(vn<<0)&((1<<(MplsLse+0))-1))
	return err
}
