package packets

const (
	ArpHrd = 16
	ArpPro = 16
	ArpHln = 8
	ArpPln = 8
	ArpOp  = 16
	ArpSha = 48
	ArpSpa = 32
	ArpTha = 48
	ArpTpa = 32
)

type Arp struct {
}

func (o Arp) Hrd(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 0); err != nil {
		return
	}
	v &= ((1 << (ArpHrd + 0)) - 1) >> 0
	return
}
func (o Arp) SetHrd(b []byte, vn uint16) error {
	v, err := ReadB16(b, 0)
	if err != nil {
		return err
	}
	err = WriteB16(b, 0, v|(vn<<0)&((1<<(ArpHrd+0))-1))
	return err
}
func (o Arp) Pro(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (ArpPro + 0)) - 1) >> 0
	return
}
func (o Arp) SetPro(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(ArpPro+0))-1))
	return err
}
func (o Arp) Hln(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 4); err != nil {
		return
	}
	v &= ((1 << (ArpHln + 0)) - 1) >> 0
	return
}
func (o Arp) SetHln(b []byte, vn uint8) error {
	v, err := ReadB8(b, 4)
	if err != nil {
		return err
	}
	err = WriteB8(b, 4, v|(vn<<0)&((1<<(ArpHln+0))-1))
	return err
}
func (o Arp) Pln(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 5); err != nil {
		return
	}
	v &= ((1 << (ArpPln + 0)) - 1) >> 0
	return
}
func (o Arp) SetPln(b []byte, vn uint8) error {
	v, err := ReadB8(b, 5)
	if err != nil {
		return err
	}
	err = WriteB8(b, 5, v|(vn<<0)&((1<<(ArpPln+0))-1))
	return err
}
func (o Arp) Op(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 6); err != nil {
		return
	}
	v &= ((1 << (ArpOp + 0)) - 1) >> 0
	return
}
func (o Arp) SetOp(b []byte, vn uint16) error {
	v, err := ReadB16(b, 6)
	if err != nil {
		return err
	}
	err = WriteB16(b, 6, v|(vn<<0)&((1<<(ArpOp+0))-1))
	return err
}
func (o Arp) Sha(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 8); err != nil {
		return
	}
	v &= ((1 << (ArpSha + 0)) - 1) >> 0
	return
}
func (o Arp) SetSha(b []byte, vn uint64) error {
	v, err := ReadB64(b, 8)
	if err != nil {
		return err
	}
	err = WriteB64(b, 8, v|(vn<<0)&((1<<(ArpSha+0))-1))
	return err
}
func (o Arp) Spa(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 14); err != nil {
		return
	}
	v &= ((1 << (ArpSpa + 0)) - 1) >> 0
	return
}
func (o Arp) SetSpa(b []byte, vn uint32) error {
	v, err := ReadB32(b, 14)
	if err != nil {
		return err
	}
	err = WriteB32(b, 14, v|(vn<<0)&((1<<(ArpSpa+0))-1))
	return err
}
func (o Arp) Tha(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 18); err != nil {
		return
	}
	v &= ((1 << (ArpTha + 0)) - 1) >> 0
	return
}
func (o Arp) SetTha(b []byte, vn uint64) error {
	v, err := ReadB64(b, 18)
	if err != nil {
		return err
	}
	err = WriteB64(b, 18, v|(vn<<0)&((1<<(ArpTha+0))-1))
	return err
}
func (o Arp) Tpa(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 24); err != nil {
		return
	}
	v &= ((1 << (ArpTpa + 0)) - 1) >> 0
	return
}
func (o Arp) SetTpa(b []byte, vn uint32) error {
	v, err := ReadB32(b, 24)
	if err != nil {
		return err
	}
	err = WriteB32(b, 24, v|(vn<<0)&((1<<(ArpTpa+0))-1))
	return err
}
