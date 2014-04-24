package packets

const (
	MplsLse = 32
)

type Mpls struct {
}

func (o Mpls) Lse(b []byte) (v uint32) {
	v, _ = ReadB32(b, 0)
	v &= ((1 << (MplsLse + 0)) - 1) >> 0
	return
}
func (o Mpls) SetLse(b []byte, vn uint32) {
	v, _ := ReadB32(b, 0)
	WriteB32(b, 0, v|(vn<<0)&((1<<(MplsLse+0))-1))
}
