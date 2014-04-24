package packets

const (
	SctpSrc  = 16
	SctpDst  = 16
	SctpVtag = 32
	SctpCsum = 32
)

type Sctp struct {
}

func (o Sctp) Src(b []byte) (v uint16) {
	v, _ = ReadB16(b, 0)
	v &= ((1 << (SctpSrc + 0)) - 1) >> 0
	return
}
func (o Sctp) SetSrc(b []byte, vn uint16) {
	v, _ := ReadB16(b, 0)
	WriteB16(b, 0, v|(vn<<0)&((1<<(SctpSrc+0))-1))
}
func (o Sctp) Dst(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (SctpDst + 0)) - 1) >> 0
	return
}
func (o Sctp) SetDst(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(SctpDst+0))-1))
}
func (o Sctp) Vtag(b []byte) (v uint32) {
	v, _ = ReadB32(b, 4)
	v &= ((1 << (SctpVtag + 0)) - 1) >> 0
	return
}
func (o Sctp) SetVtag(b []byte, vn uint32) {
	v, _ := ReadB32(b, 4)
	WriteB32(b, 4, v|(vn<<0)&((1<<(SctpVtag+0))-1))
}
func (o Sctp) Csum(b []byte) (v uint32) {
	v, _ = ReadB32(b, 8)
	v &= ((1 << (SctpCsum + 0)) - 1) >> 0
	return
}
func (o Sctp) SetCsum(b []byte, vn uint32) {
	v, _ := ReadB32(b, 8)
	WriteB32(b, 8, v|(vn<<0)&((1<<(SctpCsum+0))-1))
}
