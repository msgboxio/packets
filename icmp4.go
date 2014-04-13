package packets

const (
	Icmp4Type     = 8
	Icmp4Code     = 8
	Icmp4Checksum = 16
	Icmp4Specific = 32
)

type Icmp4 struct {
}

func (o Icmp4) Type(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Icmp4Type + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetType(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(Icmp4Type+0))-1))
}
func (o Icmp4) Code(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (Icmp4Code + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetCode(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(Icmp4Code+0))-1))
}
func (o Icmp4) Checksum(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (Icmp4Checksum + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetChecksum(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(Icmp4Checksum+0))-1))
}
func (o Icmp4) Specific(b []byte) (v uint32) {
	v, _ = ReadB32(b, 4)
	v &= ((1 << (Icmp4Specific + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetSpecific(b []byte, vn uint32) {
	v, _ := ReadB32(b, 4)
	WriteB32(b, 4, v|(vn<<0)&((1<<(Icmp4Specific+0))-1))
}
