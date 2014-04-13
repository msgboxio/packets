package packets

const (
	Ip6FragHeaderNext           = 8
	Ip6FragHeaderReserverd      = 8
	Ip6FragHeaderFragment       = 13
	Ip6FragHeaderRes            = 2
	Ip6FragHeaderMbit           = 1
	Ip6FragHeaderIdentification = 32
)

type Ip6FragHeader struct {
}

func (o Ip6FragHeader) Next(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Ip6FragHeaderNext + 0)) - 1) >> 0
	return
}
func (o Ip6FragHeader) SetNext(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(Ip6FragHeaderNext+0))-1))
}
func (o Ip6FragHeader) Reserverd(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (Ip6FragHeaderReserverd + 0)) - 1) >> 0
	return
}
func (o Ip6FragHeader) SetReserverd(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(Ip6FragHeaderReserverd+0))-1))
}
func (o Ip6FragHeader) Fragment(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (Ip6FragHeaderFragment + 0)) - 1) >> 0
	return
}
func (o Ip6FragHeader) SetFragment(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(Ip6FragHeaderFragment+0))-1))
}
func (o Ip6FragHeader) Res(b []byte) (v uint8) {
	v, _ = ReadB8(b, 3)
	v &= ((1 << (Ip6FragHeaderRes + 5)) - 1) >> 5
	return
}
func (o Ip6FragHeader) SetRes(b []byte, vn uint8) {
	v, _ := ReadB8(b, 3)
	WriteB8(b, 3, v|(vn<<5)&((1<<(Ip6FragHeaderRes+5))-1))
}
func (o Ip6FragHeader) Mbit(b []byte) (v uint8) {
	v, _ = ReadB8(b, 3)
	v &= ((1 << (Ip6FragHeaderMbit + 7)) - 1) >> 7
	return
}
func (o Ip6FragHeader) SetMbit(b []byte, vn uint8) {
	v, _ := ReadB8(b, 3)
	WriteB8(b, 3, v|(vn<<7)&((1<<(Ip6FragHeaderMbit+7))-1))
}
func (o Ip6FragHeader) Identification(b []byte) (v uint32) {
	v, _ = ReadB32(b, 4)
	v &= ((1 << (Ip6FragHeaderIdentification + 0)) - 1) >> 0
	return
}
func (o Ip6FragHeader) SetIdentification(b []byte, vn uint32) {
	v, _ := ReadB32(b, 4)
	WriteB32(b, 4, v|(vn<<0)&((1<<(Ip6FragHeaderIdentification+0))-1))
}
