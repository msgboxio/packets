package packets

const (
	LlcDsap = 8
	LlcSsap = 8
	LlcCntl = 8
)

type Llc struct {
}

func (o Llc) Dsap(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (LlcDsap + 0)) - 1) >> 0
	return
}
func (o Llc) SetDsap(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(LlcDsap+0))-1))
}
func (o Llc) Ssap(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (LlcSsap + 0)) - 1) >> 0
	return
}
func (o Llc) SetSsap(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(LlcSsap+0))-1))
}
func (o Llc) Cntl(b []byte) (v uint8) {
	v, _ = ReadB8(b, 2)
	v &= ((1 << (LlcCntl + 0)) - 1) >> 0
	return
}
func (o Llc) SetCntl(b []byte, vn uint8) {
	v, _ := ReadB8(b, 2)
	WriteB8(b, 2, v|(vn<<0)&((1<<(LlcCntl+0))-1))
}
