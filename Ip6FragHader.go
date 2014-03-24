package packets

const (
	Ip6FragHaderNext           = 8
	Ip6FragHaderReserverd      = 8
	Ip6FragHaderFragment       = 13
	Ip6FragHaderRes            = 2
	Ip6FragHaderMbit           = 1
	Ip6FragHaderIdentification = 32
)

type Ip6FragHader struct {
}

func (o Ip6FragHader) Next(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (Ip6FragHaderNext + 0)) - 1) >> 0
	return
}
func (o Ip6FragHader) SetNext(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(Ip6FragHaderNext+0))-1))
	return err
}
func (o Ip6FragHader) Reserverd(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (Ip6FragHaderReserverd + 0)) - 1) >> 0
	return
}
func (o Ip6FragHader) SetReserverd(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(Ip6FragHaderReserverd+0))-1))
	return err
}
func (o Ip6FragHader) Fragment(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (Ip6FragHaderFragment + 0)) - 1) >> 0
	return
}
func (o Ip6FragHader) SetFragment(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(Ip6FragHaderFragment+0))-1))
	return err
}
func (o Ip6FragHader) Res(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 3); err != nil {
		return
	}
	v &= ((1 << (Ip6FragHaderRes + 5)) - 1) >> 5
	return
}
func (o Ip6FragHader) SetRes(b []byte, vn uint8) error {
	v, err := ReadB8(b, 3)
	if err != nil {
		return err
	}
	err = WriteB8(b, 3, v|(vn<<5)&((1<<(Ip6FragHaderRes+5))-1))
	return err
}
func (o Ip6FragHader) Mbit(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 3); err != nil {
		return
	}
	v &= ((1 << (Ip6FragHaderMbit + 7)) - 1) >> 7
	return
}
func (o Ip6FragHader) SetMbit(b []byte, vn uint8) error {
	v, err := ReadB8(b, 3)
	if err != nil {
		return err
	}
	err = WriteB8(b, 3, v|(vn<<7)&((1<<(Ip6FragHaderMbit+7))-1))
	return err
}
func (o Ip6FragHader) Identification(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 4); err != nil {
		return
	}
	v &= ((1 << (Ip6FragHaderIdentification + 0)) - 1) >> 0
	return
}
func (o Ip6FragHader) SetIdentification(b []byte, vn uint32) error {
	v, err := ReadB32(b, 4)
	if err != nil {
		return err
	}
	err = WriteB32(b, 4, v|(vn<<0)&((1<<(Ip6FragHaderIdentification+0))-1))
	return err
}
