package packets

const (
	SnapOrg  = 24
	SnapType = 16
)

type Snap struct {
}

func (o Snap) Org(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 0); err != nil {
		return
	}
	v &= ((1 << (SnapOrg + 0)) - 1) >> 0
	return
}
func (o Snap) SetOrg(b []byte, vn uint32) error {
	v, err := ReadB32(b, 0)
	if err != nil {
		return err
	}
	err = WriteB32(b, 0, v|(vn<<0)&((1<<(SnapOrg+0))-1))
	return err
}
func (o Snap) Type(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 3); err != nil {
		return
	}
	v &= ((1 << (SnapType + 0)) - 1) >> 0
	return
}
func (o Snap) SetType(b []byte, vn uint16) error {
	v, err := ReadB16(b, 3)
	if err != nil {
		return err
	}
	err = WriteB16(b, 3, v|(vn<<0)&((1<<(SnapType+0))-1))
	return err
}
