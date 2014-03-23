package packets

const (
	ethernet_dmac = 48
	ethernet_smac = 48
	ethernet_type = 16
)

type ethernet struct {
}

func (o ethernet) Dmac(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 0); err != nil {
		return
	}
	v &= ((1 << (ethernet_dmac + 0)) - 1) >> 0
	return
}
func (o ethernet) SetDmac(b []byte, vn uint64) error {
	v, err := read_b64(b, 0)
	if err != nil {
		return err
	}
	err = write_b64(b, 0, v|(vn<<0)&((1<<(ethernet_dmac+0))-1))
	return err
}
func (o ethernet) Smac(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 6); err != nil {
		return
	}
	v &= ((1 << (ethernet_smac + 0)) - 1) >> 0
	return
}
func (o ethernet) SetSmac(b []byte, vn uint64) error {
	v, err := read_b64(b, 6)
	if err != nil {
		return err
	}
	err = write_b64(b, 6, v|(vn<<0)&((1<<(ethernet_smac+0))-1))
	return err
}
func (o ethernet) Type(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 12); err != nil {
		return
	}
	v &= ((1 << (ethernet_type + 0)) - 1) >> 0
	return
}
func (o ethernet) SetType(b []byte, vn uint16) error {
	v, err := read_b16(b, 12)
	if err != nil {
		return err
	}
	err = write_b16(b, 12, v|(vn<<0)&((1<<(ethernet_type+0))-1))
	return err
}
