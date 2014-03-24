package packets

const (
	Icmp6Type     = 8
	Icmp6Code     = 8
	Icmp6Checksum = 16
	Icmp6Specific = 32
)

type Icmp6 struct {
}

func (o Icmp6) Type(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (Icmp6Type + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetType(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(Icmp6Type+0))-1))
	return err
}
func (o Icmp6) Code(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (Icmp6Code + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetCode(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(Icmp6Code+0))-1))
	return err
}
func (o Icmp6) Checksum(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (Icmp6Checksum + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetChecksum(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(Icmp6Checksum+0))-1))
	return err
}
func (o Icmp6) Specific(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 4); err != nil {
		return
	}
	v &= ((1 << (Icmp6Specific + 0)) - 1) >> 0
	return
}
func (o Icmp6) SetSpecific(b []byte, vn uint32) error {
	v, err := ReadB32(b, 4)
	if err != nil {
		return err
	}
	err = WriteB32(b, 4, v|(vn<<0)&((1<<(Icmp6Specific+0))-1))
	return err
}
