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

func (o Tcp) Src(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 0); err != nil {
		return
	}
	v &= ((1 << (TcpSrc + 0)) - 1) >> 0
	return
}
func (o Tcp) SetSrc(b []byte, vn uint16) error {
	v, err := ReadB16(b, 0)
	if err != nil {
		return err
	}
	err = WriteB16(b, 0, v|(vn<<0)&((1<<(TcpSrc+0))-1))
	return err
}
func (o Tcp) Dst(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (TcpDst + 0)) - 1) >> 0
	return
}
func (o Tcp) SetDst(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(TcpDst+0))-1))
	return err
}
func (o Tcp) Seq(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 4); err != nil {
		return
	}
	v &= ((1 << (TcpSeq + 0)) - 1) >> 0
	return
}
func (o Tcp) SetSeq(b []byte, vn uint32) error {
	v, err := ReadB32(b, 4)
	if err != nil {
		return err
	}
	err = WriteB32(b, 4, v|(vn<<0)&((1<<(TcpSeq+0))-1))
	return err
}
func (o Tcp) Ack(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 8); err != nil {
		return
	}
	v &= ((1 << (TcpAck + 0)) - 1) >> 0
	return
}
func (o Tcp) SetAck(b []byte, vn uint32) error {
	v, err := ReadB32(b, 8)
	if err != nil {
		return err
	}
	err = WriteB32(b, 8, v|(vn<<0)&((1<<(TcpAck+0))-1))
	return err
}
func (o Tcp) Ctl(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 12); err != nil {
		return
	}
	v &= ((1 << (TcpCtl + 0)) - 1) >> 0
	return
}
func (o Tcp) SetCtl(b []byte, vn uint16) error {
	v, err := ReadB16(b, 12)
	if err != nil {
		return err
	}
	err = WriteB16(b, 12, v|(vn<<0)&((1<<(TcpCtl+0))-1))
	return err
}
func (o Tcp) Winsz(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 14); err != nil {
		return
	}
	v &= ((1 << (TcpWinsz + 0)) - 1) >> 0
	return
}
func (o Tcp) SetWinsz(b []byte, vn uint16) error {
	v, err := ReadB16(b, 14)
	if err != nil {
		return err
	}
	err = WriteB16(b, 14, v|(vn<<0)&((1<<(TcpWinsz+0))-1))
	return err
}
func (o Tcp) Csum(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 16); err != nil {
		return
	}
	v &= ((1 << (TcpCsum + 0)) - 1) >> 0
	return
}
func (o Tcp) SetCsum(b []byte, vn uint16) error {
	v, err := ReadB16(b, 16)
	if err != nil {
		return err
	}
	err = WriteB16(b, 16, v|(vn<<0)&((1<<(TcpCsum+0))-1))
	return err
}
func (o Tcp) Urg(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 18); err != nil {
		return
	}
	v &= ((1 << (TcpUrg + 0)) - 1) >> 0
	return
}
func (o Tcp) SetUrg(b []byte, vn uint16) error {
	v, err := ReadB16(b, 18)
	if err != nil {
		return err
	}
	err = WriteB16(b, 18, v|(vn<<0)&((1<<(TcpUrg+0))-1))
	return err
}
