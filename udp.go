package packets

const (
	UdpSrc  = 16
	UdpDst  = 16
	UdpLen  = 16
	UdpCsum = 16
)

type Udp struct {
}

func (o Udp) Src(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 0); err != nil {
		return
	}
	v &= ((1 << (UdpSrc + 0)) - 1) >> 0
	return
}
func (o Udp) SetSrc(b []byte, vn uint16) error {
	v, err := ReadB16(b, 0)
	if err != nil {
		return err
	}
	err = WriteB16(b, 0, v|(vn<<0)&((1<<(UdpSrc+0))-1))
	return err
}
func (o Udp) Dst(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (UdpDst + 0)) - 1) >> 0
	return
}
func (o Udp) SetDst(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(UdpDst+0))-1))
	return err
}
func (o Udp) Len(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 4); err != nil {
		return
	}
	v &= ((1 << (UdpLen + 0)) - 1) >> 0
	return
}
func (o Udp) SetLen(b []byte, vn uint16) error {
	v, err := ReadB16(b, 4)
	if err != nil {
		return err
	}
	err = WriteB16(b, 4, v|(vn<<0)&((1<<(UdpLen+0))-1))
	return err
}
func (o Udp) Csum(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 6); err != nil {
		return
	}
	v &= ((1 << (UdpCsum + 0)) - 1) >> 0
	return
}
func (o Udp) SetCsum(b []byte, vn uint16) error {
	v, err := ReadB16(b, 6)
	if err != nil {
		return err
	}
	err = WriteB16(b, 6, v|(vn<<0)&((1<<(UdpCsum+0))-1))
	return err
}
