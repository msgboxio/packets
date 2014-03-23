package packets

const (
	arp_hrd = 16
	arp_pro = 16
	arp_hln = 8
	arp_pln = 8
	arp_op  = 16
	arp_sha = 48
	arp_spa = 32
	arp_tha = 48
	arp_tpa = 32
)

type arp struct {
}

func (o arp) Hrd(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 0); err != nil {
		return
	}
	v &= ((1 << (arp_hrd + 0)) - 1) >> 0
	return
}
func (o arp) SetHrd(b []byte, vn uint16) error {
	v, err := read_b16(b, 0)
	if err != nil {
		return err
	}
	err = write_b16(b, 0, v|(vn<<0)&((1<<(arp_hrd+0))-1))
	return err
}
func (o arp) Pro(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (arp_pro + 0)) - 1) >> 0
	return
}
func (o arp) SetPro(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(arp_pro+0))-1))
	return err
}
func (o arp) Hln(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 4); err != nil {
		return
	}
	v &= ((1 << (arp_hln + 0)) - 1) >> 0
	return
}
func (o arp) SetHln(b []byte, vn uint8) error {
	v, err := read_b8(b, 4)
	if err != nil {
		return err
	}
	err = write_b8(b, 4, v|(vn<<0)&((1<<(arp_hln+0))-1))
	return err
}
func (o arp) Pln(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 5); err != nil {
		return
	}
	v &= ((1 << (arp_pln + 0)) - 1) >> 0
	return
}
func (o arp) SetPln(b []byte, vn uint8) error {
	v, err := read_b8(b, 5)
	if err != nil {
		return err
	}
	err = write_b8(b, 5, v|(vn<<0)&((1<<(arp_pln+0))-1))
	return err
}
func (o arp) Op(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 6); err != nil {
		return
	}
	v &= ((1 << (arp_op + 0)) - 1) >> 0
	return
}
func (o arp) SetOp(b []byte, vn uint16) error {
	v, err := read_b16(b, 6)
	if err != nil {
		return err
	}
	err = write_b16(b, 6, v|(vn<<0)&((1<<(arp_op+0))-1))
	return err
}
func (o arp) Sha(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 8); err != nil {
		return
	}
	v &= ((1 << (arp_sha + 0)) - 1) >> 0
	return
}
func (o arp) SetSha(b []byte, vn uint64) error {
	v, err := read_b64(b, 8)
	if err != nil {
		return err
	}
	err = write_b64(b, 8, v|(vn<<0)&((1<<(arp_sha+0))-1))
	return err
}
func (o arp) Spa(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 14); err != nil {
		return
	}
	v &= ((1 << (arp_spa + 0)) - 1) >> 0
	return
}
func (o arp) SetSpa(b []byte, vn uint32) error {
	v, err := read_b32(b, 14)
	if err != nil {
		return err
	}
	err = write_b32(b, 14, v|(vn<<0)&((1<<(arp_spa+0))-1))
	return err
}
func (o arp) Tha(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 18); err != nil {
		return
	}
	v &= ((1 << (arp_tha + 0)) - 1) >> 0
	return
}
func (o arp) SetTha(b []byte, vn uint64) error {
	v, err := read_b64(b, 18)
	if err != nil {
		return err
	}
	err = write_b64(b, 18, v|(vn<<0)&((1<<(arp_tha+0))-1))
	return err
}
func (o arp) Tpa(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 24); err != nil {
		return
	}
	v &= ((1 << (arp_tpa + 0)) - 1) >> 0
	return
}
func (o arp) SetTpa(b []byte, vn uint32) error {
	v, err := read_b32(b, 24)
	if err != nil {
		return err
	}
	err = write_b32(b, 24, v|(vn<<0)&((1<<(arp_tpa+0))-1))
	return err
}
