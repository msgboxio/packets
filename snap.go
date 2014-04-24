package packets

const (
	SnapOrg  = 24
	SnapType = 16
)

type Snap struct {
}

func (o Snap) Org(b []byte) (v uint32) {
	v, _ = ReadB32(b, 0)
	v &= ((1 << (SnapOrg + 0)) - 1) >> 0
	return
}
func (o Snap) SetOrg(b []byte, vn uint32) {
	v, _ := ReadB32(b, 0)
	WriteB32(b, 0, v|(vn<<0)&((1<<(SnapOrg+0))-1))
}
func (o Snap) Type(b []byte) (v uint16) {
	v, _ = ReadB16(b, 3)
	v &= ((1 << (SnapType + 0)) - 1) >> 0
	return
}
func (o Snap) SetType(b []byte, vn uint16) {
	v, _ := ReadB16(b, 3)
	WriteB16(b, 3, v|(vn<<0)&((1<<(SnapType+0))-1))
}
