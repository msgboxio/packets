package packets

const (
	Icmp6Type     = 8
	Icmp6Code     = 8
	Icmp6Checksum = 16
	Icmp6Specific = 32
)

type Icmp6 struct {
}

func (o Icmp6) Type(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Icmp6Type + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetType(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(Icmp6Type+0))-1))
}
func (o Icmp6) Code(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (Icmp6Code + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetCode(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(Icmp6Code+0))-1))
}
func (o Icmp6) Checksum(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (Icmp6Checksum + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetChecksum(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(Icmp6Checksum+0))-1))
}
func (o Icmp6) Specific(b []byte) (v uint32) {
	v, _ = ReadB32(b, 4)
	v &= ((1 << (Icmp6Specific + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetSpecific(b []byte, vn uint32) {
	v, _ := ReadB32(b, 4)
	WriteB32(b, 4, v|(vn<<0)&((1<<(Icmp6Specific+0))-1))
}
