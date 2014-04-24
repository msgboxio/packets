package packets

const (
	Ip6Version  = 4
	Ip6Tc1      = 4
	Ip6Tc2      = 4
	Ip6Flow     = 20
	Ip6Payload  = 16
	Ip6Next     = 8
	Ip6HopLimit = 8
	Ip6Source   = 128
	Ip6Dest     = 128
)

type Ip6 struct {
}

func (o Ip6) Version(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Ip6Version + 0)) - 1) >> 0
	return
}
func (o Ip6) SetVersion(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(Ip6Version+0))-1))
}
func (o Ip6) Tc1(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Ip6Tc1 + 4)) - 1) >> 4
	return
}
func (o Ip6) SetTc1(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<4)&((1<<(Ip6Tc1+4))-1))
}
func (o Ip6) Tc2(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (Ip6Tc2 + 0)) - 1) >> 0
	return
}
func (o Ip6) SetTc2(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(Ip6Tc2+0))-1))
}
func (o Ip6) Flow(b []byte) (v uint32) {
	v, _ = ReadB32(b, 1)
	v &= ((1 << (Ip6Flow + 4)) - 1) >> 4
	return
}
func (o Ip6) SetFlow(b []byte, vn uint32) {
	v, _ := ReadB32(b, 1)
	WriteB32(b, 1, v|(vn<<4)&((1<<(Ip6Flow+4))-1))
}
func (o Ip6) Payload(b []byte) (v uint16) {
	v, _ = ReadB16(b, 4)
	v &= ((1 << (Ip6Payload + 0)) - 1) >> 0
	return
}
func (o Ip6) SetPayload(b []byte, vn uint16) {
	v, _ := ReadB16(b, 4)
	WriteB16(b, 4, v|(vn<<0)&((1<<(Ip6Payload+0))-1))
}
func (o Ip6) Next(b []byte) (v uint8) {
	v, _ = ReadB8(b, 6)
	v &= ((1 << (Ip6Next + 0)) - 1) >> 0
	return
}
func (o Ip6) SetNext(b []byte, vn uint8) {
	v, _ := ReadB8(b, 6)
	WriteB8(b, 6, v|(vn<<0)&((1<<(Ip6Next+0))-1))
}
func (o Ip6) HopLimit(b []byte) (v uint8) {
	v, _ = ReadB8(b, 7)
	v &= ((1 << (Ip6HopLimit + 0)) - 1) >> 0
	return
}
func (o Ip6) SetHopLimit(b []byte, vn uint8) {
	v, _ := ReadB8(b, 7)
	WriteB8(b, 7, v|(vn<<0)&((1<<(Ip6HopLimit+0))-1))
}
func (o Ip6) Source(b []byte) []byte {
	return b[8:24]
}
func (o Ip6) Dest(b []byte) []byte {
	return b[24:40]
}
