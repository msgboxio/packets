package packets

const (
	Icmp4Type     = 8
	Icmp4Code     = 8
	Icmp4Checksum = 16
	Icmp4Specific = 32
)

type Icmp4 struct {
}

func (o Icmp4) Type(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (Icmp4Type + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetType(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(Icmp4Type+0))-1))
	return err
}
func (o Icmp4) Code(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (Icmp4Code + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetCode(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(Icmp4Code+0))-1))
	return err
}
func (o Icmp4) Checksum(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (Icmp4Checksum + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetChecksum(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(Icmp4Checksum+0))-1))
	return err
}
func (o Icmp4) Specific(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 4); err != nil {
		return
	}
	v &= ((1 << (Icmp4Specific + 0)) - 1) >> 0
	return
}
func (o Icmp4) SetSpecific(b []byte, vn uint32) error {
	v, err := ReadB32(b, 4)
	if err != nil {
		return err
	}
	err = WriteB32(b, 4, v|(vn<<0)&((1<<(Icmp4Specific+0))-1))
	return err
}
