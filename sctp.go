package packets

const (
	sctp_src  = 16
	sctp_dst  = 16
	sctp_vtag = 32
	sctp_csum = 32
)

type sctp struct {
}

func (o sctp) Src(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 0); err != nil {
		return
	}
	v &= ((1 << (sctp_src + 0)) - 1) >> 0
	return
}
func (o sctp) SetSrc(b []byte, vn uint16) error {
	v, err := read_b16(b, 0)
	if err != nil {
		return err
	}
	err = write_b16(b, 0, v|(vn<<0)&((1<<(sctp_src+0))-1))
	return err
}
func (o sctp) Dst(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (sctp_dst + 0)) - 1) >> 0
	return
}
func (o sctp) SetDst(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(sctp_dst+0))-1))
	return err
}
func (o sctp) Vtag(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 4); err != nil {
		return
	}
	v &= ((1 << (sctp_vtag + 0)) - 1) >> 0
	return
}
func (o sctp) SetVtag(b []byte, vn uint32) error {
	v, err := read_b32(b, 4)
	if err != nil {
		return err
	}
	err = write_b32(b, 4, v|(vn<<0)&((1<<(sctp_vtag+0))-1))
	return err
}
func (o sctp) Csum(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 8); err != nil {
		return
	}
	v &= ((1 << (sctp_csum + 0)) - 1) >> 0
	return
}
func (o sctp) SetCsum(b []byte, vn uint32) error {
	v, err := read_b32(b, 8)
	if err != nil {
		return err
	}
	err = write_b32(b, 8, v|(vn<<0)&((1<<(sctp_csum+0))-1))
	return err
}
