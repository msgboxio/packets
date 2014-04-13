package packets

const (
	NdOptNdtype = 8
	NdOptLength = 8
)

type NdOpt struct {
}

func (o NdOpt) Ndtype(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (NdOptNdtype + 0)) - 1) >> 0
	return
}
func (o NdOpt) SetNdtype(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(NdOptNdtype+0))-1))
}
func (o NdOpt) Length(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (NdOptLength + 0)) - 1) >> 0
	return
}
func (o NdOpt) SetLength(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(NdOptLength+0))-1))
}
