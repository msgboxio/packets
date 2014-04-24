package packets

const (
	VlanQtagTpid = 16
	VlanQtagTci  = 16
)

type VlanQtag struct {
}

func (o VlanQtag) Tpid(b []byte) (v uint16) {
	v, _ = ReadB16(b, 0)
	v &= ((1 << (VlanQtagTpid + 0)) - 1) >> 0
	return
}
func (o VlanQtag) SetTpid(b []byte, vn uint16) {
	v, _ := ReadB16(b, 0)
	WriteB16(b, 0, v|(vn<<0)&((1<<(VlanQtagTpid+0))-1))
}
func (o VlanQtag) Tci(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (VlanQtagTci + 0)) - 1) >> 0
	return
}
func (o VlanQtag) SetTci(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(VlanQtagTci+0))-1))
}
