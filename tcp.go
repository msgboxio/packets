package packets

const (
	TcpSrc   = 16
	TcpDst   = 16
	TcpSeq   = 32
	TcpAck   = 32
	TcpCtl   = 16
	TcpWinsz = 16
	TcpCsum  = 16
	TcpUrg   = 16
)

type Tcp struct {
}

func (o Tcp) Src(b []byte) (v uint16) {
	v, _ = ReadB16(b, 0)
	v &= ((1 << (TcpSrc + 0)) - 1) >> 0
	return
}
func (o Tcp) SetSrc(b []byte, vn uint16) {
	v, _ := ReadB16(b, 0)
	WriteB16(b, 0, v|(vn<<0)&((1<<(TcpSrc+0))-1))
}
func (o Tcp) Dst(b []byte) (v uint16) {
	v, _ = ReadB16(b, 2)
	v &= ((1 << (TcpDst + 0)) - 1) >> 0
	return
}
func (o Tcp) SetDst(b []byte, vn uint16) {
	v, _ := ReadB16(b, 2)
	WriteB16(b, 2, v|(vn<<0)&((1<<(TcpDst+0))-1))
}
func (o Tcp) Seq(b []byte) (v uint32) {
	v, _ = ReadB32(b, 4)
	v &= ((1 << (TcpSeq + 0)) - 1) >> 0
	return
}
func (o Tcp) SetSeq(b []byte, vn uint32) {
	v, _ := ReadB32(b, 4)
	WriteB32(b, 4, v|(vn<<0)&((1<<(TcpSeq+0))-1))
}
func (o Tcp) Ack(b []byte) (v uint32) {
	v, _ = ReadB32(b, 8)
	v &= ((1 << (TcpAck + 0)) - 1) >> 0
	return
}
func (o Tcp) SetAck(b []byte, vn uint32) {
	v, _ := ReadB32(b, 8)
	WriteB32(b, 8, v|(vn<<0)&((1<<(TcpAck+0))-1))
}
func (o Tcp) Ctl(b []byte) (v uint16) {
	v, _ = ReadB16(b, 12)
	v &= ((1 << (TcpCtl + 0)) - 1) >> 0
	return
}
func (o Tcp) SetCtl(b []byte, vn uint16) {
	v, _ := ReadB16(b, 12)
	WriteB16(b, 12, v|(vn<<0)&((1<<(TcpCtl+0))-1))
}
func (o Tcp) Winsz(b []byte) (v uint16) {
	v, _ = ReadB16(b, 14)
	v &= ((1 << (TcpWinsz + 0)) - 1) >> 0
	return
}
func (o Tcp) SetWinsz(b []byte, vn uint16) {
	v, _ := ReadB16(b, 14)
	WriteB16(b, 14, v|(vn<<0)&((1<<(TcpWinsz+0))-1))
}
func (o Tcp) Csum(b []byte) (v uint16) {
	v, _ = ReadB16(b, 16)
	v &= ((1 << (TcpCsum + 0)) - 1) >> 0
	return
}
func (o Tcp) SetCsum(b []byte, vn uint16) {
	v, _ := ReadB16(b, 16)
	WriteB16(b, 16, v|(vn<<0)&((1<<(TcpCsum+0))-1))
}
func (o Tcp) Urg(b []byte) (v uint16) {
	v, _ = ReadB16(b, 18)
	v &= ((1 << (TcpUrg + 0)) - 1) >> 0
	return
}
func (o Tcp) SetUrg(b []byte, vn uint16) {
	v, _ := ReadB16(b, 18)
	WriteB16(b, 18, v|(vn<<0)&((1<<(TcpUrg+0))-1))
}
