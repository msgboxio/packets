package packets

const (
	Ip6ExtNext = 8
	Ip6ExtLen  = 8
)

type Ip6Ext struct {
}

func (o Ip6Ext) Next(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Ip6ExtNext + 0)) - 1) >> 0
	return
}
func (o Ip6Ext) SetNext(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(Ip6ExtNext+0))-1))
}
func (o Ip6Ext) Len(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (Ip6ExtLen + 0)) - 1) >> 0
	return
}
func (o Ip6Ext) SetLen(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(Ip6ExtLen+0))-1))
}
