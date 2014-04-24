package packets

const (
	Ip4Version  = 4
	Ip4Hdrlen   = 4
	Ip4Tos      = 8
	Ip4Length   = 16
	Ip4Id       = 16
	Ip4Flags    = 3
	Ip4Offset   = 13
	Ip4Ttl      = 8
	Ip4Protocol = 8
	Ip4Hdrchks  = 16
	Ip4Srcaddr  = 32
	Ip4Dstaddr  = 32
	Ip4Options  = '*'
)

type Ip4 struct {
}

func (o Ip4) Version(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Ip4Version + 0)) - 1) >> 0
	return
}
func (o Ip4) SetVersion(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<0)&((1<<(Ip4Version+0))-1))
}
func (o Ip4) Hdrlen(b []byte) (v uint8) {
	v, _ = ReadB8(b, 0)
	v &= ((1 << (Ip4Hdrlen + 4)) - 1) >> 4
	return
}
func (o Ip4) SetHdrlen(b []byte, vn uint8) {
	v, _ := ReadB8(b, 0)
	WriteB8(b, 0, v|(vn<<4)&((1<<(Ip4Hdrlen+4))-1))
}
func (o Ip4) Tos(b []byte) (v uint8) {
	v, _ = ReadB8(b, 1)
	v &= ((1 << (Ip4Tos + 0)) - 1) >> 0
	return
}
func (o Ip4) SetTos(b []byte, vn uint8) {
	v, _ := ReadB8(b, 1)
	WriteB8(b, 1, v|(vn<<0)&((1<<(Ip4Tos+0))-1))
}
func (o Ip4) Length(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (Ip4Length + 0)) - 1) >> 0
	return
}
func (o Ip4) SetLength(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(Ip4Length+0))-1))
}
func (o Ip4) Id(b []byte) (v uint16) {
	v, _ = ReadB16(b, 4)
	v &= ((1 << (Ip4Id + 0)) - 1) >> 0
	return
}
func (o Ip4) SetId(b []byte, vn uint16) {
	v, _ := ReadB16(b, 4)
	WriteB16(b, 4, v|(vn<<0)&((1<<(Ip4Id+0))-1))
}
func (o Ip4) Flags(b []byte) (v uint8) {
	v, _ = ReadB8(b, 6)
	v &= ((1 << (Ip4Flags + 0)) - 1) >> 0
	return
}
func (o Ip4) SetFlags(b []byte, vn uint8) {
	v, _ := ReadB8(b, 6)
	WriteB8(b, 6, v|(vn<<0)&((1<<(Ip4Flags+0))-1))
}
func (o Ip4) Offset(b []byte) (v uint16) {
	v, _ = ReadB16(b, 6)
	v &= ((1 << (Ip4Offset + 3)) - 1) >> 3
	return
}
func (o Ip4) SetOffset(b []byte, vn uint16) {
	v, _ := ReadB16(b, 6)
	WriteB16(b, 6, v|(vn<<3)&((1<<(Ip4Offset+3))-1))
}
func (o Ip4) Ttl(b []byte) (v uint8) {
	v, _ = ReadB8(b, 8)
	v &= ((1 << (Ip4Ttl + 0)) - 1) >> 0
	return
}
func (o Ip4) SetTtl(b []byte, vn uint8) {
	v, _ := ReadB8(b, 8)
	WriteB8(b, 8, v|(vn<<0)&((1<<(Ip4Ttl+0))-1))
}
func (o Ip4) Protocol(b []byte) (v uint8) {
	v, _ = ReadB8(b, 9)
	v &= ((1 << (Ip4Protocol + 0)) - 1) >> 0
	return
}
func (o Ip4) SetProtocol(b []byte, vn uint8) {
	v, _ := ReadB8(b, 9)
	WriteB8(b, 9, v|(vn<<0)&((1<<(Ip4Protocol+0))-1))
}
func (o Ip4) Hdrchks(b []byte) (v uint16) {
	v, _ = ReadB16(b, 10)
	v &= ((1 << (Ip4Hdrchks + 0)) - 1) >> 0
	return
}
func (o Ip4) SetHdrchks(b []byte, vn uint16) {
	v, _ := ReadB16(b, 10)
	WriteB16(b, 10, v|(vn<<0)&((1<<(Ip4Hdrchks+0))-1))
}
func (o Ip4) Srcaddr(b []byte) (v uint32) {
	v, _ = ReadB32(b, 12)
	v &= ((1 << (Ip4Srcaddr + 0)) - 1) >> 0
	return
}
func (o Ip4) SetSrcaddr(b []byte, vn uint32) {
	v, _ := ReadB32(b, 12)
	WriteB32(b, 12, v|(vn<<0)&((1<<(Ip4Srcaddr+0))-1))
}
func (o Ip4) Dstaddr(b []byte) (v uint32) {
	v, _ = ReadB32(b, 16)
	v &= ((1 << (Ip4Dstaddr + 0)) - 1) >> 0
	return
}
func (o Ip4) SetDstaddr(b []byte, vn uint32) {
	v, _ := ReadB32(b, 16)
	WriteB32(b, 16, v|(vn<<0)&((1<<(Ip4Dstaddr+0))-1))
}
