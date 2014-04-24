package packets

const (
	ArpHrd = 16
	ArpPro = 16
	ArpHln = 8
	ArpPln = 8
	ArpOp  = 16
	ArpSha = 48
	ArpSpa = 32
	ArpTha = 48
	ArpTpa = 32
)

type Arp struct {
}

func (o Arp) Hrd(b []byte) (v uint16) {
	v, _ = ReadB16(b, 0)
	v &= ((1 << (ArpHrd + 0)) - 1) >> 0
	return
}
func (o Arp) SetHrd(b []byte, vn uint16) {
	v, _ := ReadB16(b, 0)
	WriteB16(b, 0, v|(vn<<0)&((1<<(ArpHrd+0))-1))
}
func (o Arp) Pro(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (ArpPro + 0)) - 1) >> 0
	return
}
func (o Arp) SetPro(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(ArpPro+0))-1))
}
func (o Arp) Hln(b []byte) (v uint8) {
	v, _ = ReadB8(b, 4)
	v &= ((1 << (ArpHln + 0)) - 1) >> 0
	return
}
func (o Arp) SetHln(b []byte, vn uint8) {
	v, _ := ReadB8(b, 4)
	WriteB8(b, 4, v|(vn<<0)&((1<<(ArpHln+0))-1))
}
func (o Arp) Pln(b []byte) (v uint8) {
	v, _ = ReadB8(b, 5)
	v &= ((1 << (ArpPln + 0)) - 1) >> 0
	return
}
func (o Arp) SetPln(b []byte, vn uint8) {
	v, _ := ReadB8(b, 5)
	WriteB8(b, 5, v|(vn<<0)&((1<<(ArpPln+0))-1))
}
func (o Arp) Op(b []byte) (v uint16) {
	v, _ = ReadB16(b, 6)
	v &= ((1 << (ArpOp + 0)) - 1) >> 0
	return
}
func (o Arp) SetOp(b []byte, vn uint16) {
	v, _ := ReadB16(b, 6)
	WriteB16(b, 6, v|(vn<<0)&((1<<(ArpOp+0))-1))
}
func (o Arp) Sha(b []byte) []byte {
	return b[8:14]
}
func (o Arp) Spa(b []byte) (v uint32) {
	v, _ = ReadB32(b, 14)
	v &= ((1 << (ArpSpa + 0)) - 1) >> 0
	return
}
func (o Arp) SetSpa(b []byte, vn uint32) {
	v, _ := ReadB32(b, 14)
	WriteB32(b, 14, v|(vn<<0)&((1<<(ArpSpa+0))-1))
}
func (o Arp) Tha(b []byte) []byte {
	return b[18:24]
}
func (o Arp) Tpa(b []byte) (v uint32) {
	v, _ = ReadB32(b, 24)
	v &= ((1 << (ArpTpa + 0)) - 1) >> 0
	return
}
func (o Arp) SetTpa(b []byte, vn uint32) {
	v, _ := ReadB32(b, 24)
	WriteB32(b, 24, v|(vn<<0)&((1<<(ArpTpa+0))-1))
}
