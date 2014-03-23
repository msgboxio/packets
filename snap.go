package packets

const (
	snap_org  = 24
	snap_type = 16
)

type snap struct {
}

func (o snap) Org(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 0); err != nil {
		return
	}
	v &= ((1 << (snap_org + 0)) - 1) >> 0
	return
}
func (o snap) SetOrg(b []byte, vn uint32) error {
	v, err := read_b32(b, 0)
	if err != nil {
		return err
	}
	err = write_b32(b, 0, v|(vn<<0)&((1<<(snap_org+0))-1))
	return err
}
func (o snap) Type(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 3); err != nil {
		return
	}
	v &= ((1 << (snap_type + 0)) - 1) >> 0
	return
}
func (o snap) SetType(b []byte, vn uint16) error {
	v, err := read_b16(b, 3)
	if err != nil {
		return err
	}
	err = write_b16(b, 3, v|(vn<<0)&((1<<(snap_type+0))-1))
	return err
}
