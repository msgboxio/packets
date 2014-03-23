package packets

const (
	ip6_frag_hader_next           = 8
	ip6_frag_hader_reserverd      = 8
	ip6_frag_hader_fragment       = 13
	ip6_frag_hader_res            = 2
	ip6_frag_hader_mbit           = 1
	ip6_frag_hader_identification = 32
)

type ip6_frag_hader struct {
}

func (o ip6_frag_hader) Next(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (ip6_frag_hader_next + 0)) - 1) >> 0
	return
}
func (o ip6_frag_hader) SetNext(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(ip6_frag_hader_next+0))-1))
	return err
}
func (o ip6_frag_hader) Reserverd(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (ip6_frag_hader_reserverd + 0)) - 1) >> 0
	return
}
func (o ip6_frag_hader) SetReserverd(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(ip6_frag_hader_reserverd+0))-1))
	return err
}
func (o ip6_frag_hader) Fragment(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (ip6_frag_hader_fragment + 0)) - 1) >> 0
	return
}
func (o ip6_frag_hader) SetFragment(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(ip6_frag_hader_fragment+0))-1))
	return err
}
func (o ip6_frag_hader) Res(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 3); err != nil {
		return
	}
	v &= ((1 << (ip6_frag_hader_res + 5)) - 1) >> 5
	return
}
func (o ip6_frag_hader) SetRes(b []byte, vn uint8) error {
	v, err := read_b8(b, 3)
	if err != nil {
		return err
	}
	err = write_b8(b, 3, v|(vn<<5)&((1<<(ip6_frag_hader_res+5))-1))
	return err
}
func (o ip6_frag_hader) Mbit(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 3); err != nil {
		return
	}
	v &= ((1 << (ip6_frag_hader_mbit + 7)) - 1) >> 7
	return
}
func (o ip6_frag_hader) SetMbit(b []byte, vn uint8) error {
	v, err := read_b8(b, 3)
	if err != nil {
		return err
	}
	err = write_b8(b, 3, v|(vn<<7)&((1<<(ip6_frag_hader_mbit+7))-1))
	return err
}
func (o ip6_frag_hader) Identification(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 4); err != nil {
		return
	}
	v &= ((1 << (ip6_frag_hader_identification + 0)) - 1) >> 0
	return
}
func (o ip6_frag_hader) SetIdentification(b []byte, vn uint32) error {
	v, err := read_b32(b, 4)
	if err != nil {
		return err
	}
	err = write_b32(b, 4, v|(vn<<0)&((1<<(ip6_frag_hader_identification+0))-1))
	return err
}
