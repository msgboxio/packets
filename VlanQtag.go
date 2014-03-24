package packets

const (
	VlanQtagTpid = 16
	VlanQtagTci  = 16
)

type VlanQtag struct {
}

func (o VlanQtag) Tpid(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 0); err != nil {
		return
	}
	v &= ((1 << (VlanQtagTpid + 0)) - 1) >> 0
	return
}
func (o VlanQtag) SetTpid(b []byte, vn uint16) error {
	v, err := ReadB16(b, 0)
	if err != nil {
		return err
	}
	err = WriteB16(b, 0, v|(vn<<0)&((1<<(VlanQtagTpid+0))-1))
	return err
}
func (o VlanQtag) Tci(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (VlanQtagTci + 0)) - 1) >> 0
	return
}
func (o VlanQtag) SetTci(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(VlanQtagTci+0))-1))
	return err
}
