package packets

const (
	EthernetDmac = 48
	EthernetSmac = 48
	EthernetType = 16
)

type Ethernet struct {
}

func (o Ethernet) Dmac(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 0); err != nil {
		return
	}
	v &= ((1 << (EthernetDmac + 0)) - 1) >> 0
	return
}
func (o Ethernet) SetDmac(b []byte, vn uint64) error {
	v, err := ReadB64(b, 0)
	if err != nil {
		return err
	}
	err = WriteB64(b, 0, v|(vn<<0)&((1<<(EthernetDmac+0))-1))
	return err
}
func (o Ethernet) Smac(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 6); err != nil {
		return
	}
	v &= ((1 << (EthernetSmac + 0)) - 1) >> 0
	return
}
func (o Ethernet) SetSmac(b []byte, vn uint64) error {
	v, err := ReadB64(b, 6)
	if err != nil {
		return err
	}
	err = WriteB64(b, 6, v|(vn<<0)&((1<<(EthernetSmac+0))-1))
	return err
}
func (o Ethernet) Type(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 12); err != nil {
		return
	}
	v &= ((1 << (EthernetType + 0)) - 1) >> 0
	return
}
func (o Ethernet) SetType(b []byte, vn uint16) error {
	v, err := ReadB16(b, 12)
	if err != nil {
		return err
	}
	err = WriteB16(b, 12, v|(vn<<0)&((1<<(EthernetType+0))-1))
	return err
}
