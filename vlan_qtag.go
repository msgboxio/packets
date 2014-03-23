package packets

const (
	vlan_qtag_tpid = 16
	vlan_qtag_tci  = 16
)

type vlan_qtag struct {
}

func (o vlan_qtag) Tpid(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 0); err != nil {
		return
	}
	v &= ((1 << (vlan_qtag_tpid + 0)) - 1) >> 0
	return
}
func (o vlan_qtag) SetTpid(b []byte, vn uint16) error {
	v, err := read_b16(b, 0)
	if err != nil {
		return err
	}
	err = write_b16(b, 0, v|(vn<<0)&((1<<(vlan_qtag_tpid+0))-1))
	return err
}
func (o vlan_qtag) Tci(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (vlan_qtag_tci + 0)) - 1) >> 0
	return
}
func (o vlan_qtag) SetTci(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(vlan_qtag_tci+0))-1))
	return err
}
