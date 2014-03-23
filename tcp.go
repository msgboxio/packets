package packets

const (
	tcp_src   = 16
	tcp_dst   = 16
	tcp_seq   = 32
	tcp_ack   = 32
	tcp_ctl   = 16
	tcp_winsz = 16
	tcp_csum  = 16
	tcp_urg   = 16
)

type tcp struct {
}

func (o tcp) Src(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 0); err != nil {
		return
	}
	v &= ((1 << (tcp_src + 0)) - 1) >> 0
	return
}
func (o tcp) SetSrc(b []byte, vn uint16) error {
	v, err := read_b16(b, 0)
	if err != nil {
		return err
	}
	err = write_b16(b, 0, v|(vn<<0)&((1<<(tcp_src+0))-1))
	return err
}
func (o tcp) Dst(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (tcp_dst + 0)) - 1) >> 0
	return
}
func (o tcp) SetDst(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(tcp_dst+0))-1))
	return err
}
func (o tcp) Seq(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 4); err != nil {
		return
	}
	v &= ((1 << (tcp_seq + 0)) - 1) >> 0
	return
}
func (o tcp) SetSeq(b []byte, vn uint32) error {
	v, err := read_b32(b, 4)
	if err != nil {
		return err
	}
	err = write_b32(b, 4, v|(vn<<0)&((1<<(tcp_seq+0))-1))
	return err
}
func (o tcp) Ack(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 8); err != nil {
		return
	}
	v &= ((1 << (tcp_ack + 0)) - 1) >> 0
	return
}
func (o tcp) SetAck(b []byte, vn uint32) error {
	v, err := read_b32(b, 8)
	if err != nil {
		return err
	}
	err = write_b32(b, 8, v|(vn<<0)&((1<<(tcp_ack+0))-1))
	return err
}
func (o tcp) Ctl(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 12); err != nil {
		return
	}
	v &= ((1 << (tcp_ctl + 0)) - 1) >> 0
	return
}
func (o tcp) SetCtl(b []byte, vn uint16) error {
	v, err := read_b16(b, 12)
	if err != nil {
		return err
	}
	err = write_b16(b, 12, v|(vn<<0)&((1<<(tcp_ctl+0))-1))
	return err
}
func (o tcp) Winsz(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 14); err != nil {
		return
	}
	v &= ((1 << (tcp_winsz + 0)) - 1) >> 0
	return
}
func (o tcp) SetWinsz(b []byte, vn uint16) error {
	v, err := read_b16(b, 14)
	if err != nil {
		return err
	}
	err = write_b16(b, 14, v|(vn<<0)&((1<<(tcp_winsz+0))-1))
	return err
}
func (o tcp) Csum(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 16); err != nil {
		return
	}
	v &= ((1 << (tcp_csum + 0)) - 1) >> 0
	return
}
func (o tcp) SetCsum(b []byte, vn uint16) error {
	v, err := read_b16(b, 16)
	if err != nil {
		return err
	}
	err = write_b16(b, 16, v|(vn<<0)&((1<<(tcp_csum+0))-1))
	return err
}
func (o tcp) Urg(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 18); err != nil {
		return
	}
	v &= ((1 << (tcp_urg + 0)) - 1) >> 0
	return
}
func (o tcp) SetUrg(b []byte, vn uint16) error {
	v, err := read_b16(b, 18)
	if err != nil {
		return err
	}
	err = write_b16(b, 18, v|(vn<<0)&((1<<(tcp_urg+0))-1))
	return err
}
