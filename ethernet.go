package packets

const (
	EthernetDmac = 48
	EthernetSmac = 48
	EthernetType = 16
)

type Ethernet struct {
}

func (o Ethernet) Dmac(b []byte) []byte {
	return b[0:6]
}
func (o Ethernet) Smac(b []byte) []byte {
	return b[6:12]
}
func (o Ethernet) Type(b []byte) (v uint16) {
	v, _ = ReadB16(b, 12)
	v &= ((1 << (EthernetType + 0)) - 1) >> 0
	return
}
func (o Ethernet) SetType(b []byte, vn uint16) {
	v, _ := ReadB16(b, 12)
	WriteB16(b, 12, v|(vn<<0)&((1<<(EthernetType+0))-1))
}
