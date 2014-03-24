package packets

const (
	Ip6Version  = 4
	Ip6Tc1      = 4
	Ip6Tc2      = 4
	Ip6Flow     = 20
	Ip6Payload  = 16
	Ip6Next     = 8
	Ip6HopLimit = 8
	Ip6Source1  = 64
	Ip6Source2  = 64
	Ip6Dest1    = 64
	Ip6Dest2    = 64
)

type Ip6 struct {
}

func (o Ip6) Version(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (Ip6Version + 0)) - 1) >> 0
	return
}
func (o Ip6) SetVersion(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(Ip6Version+0))-1))
	return err
}
func (o Ip6) Tc1(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (Ip6Tc1 + 4)) - 1) >> 4
	return
}
func (o Ip6) SetTc1(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<4)&((1<<(Ip6Tc1+4))-1))
	return err
}
func (o Ip6) Tc2(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (Ip6Tc2 + 0)) - 1) >> 0
	return
}
func (o Ip6) SetTc2(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(Ip6Tc2+0))-1))
	return err
}
func (o Ip6) Flow(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 1); err != nil {
		return
	}
	v &= ((1 << (Ip6Flow + 4)) - 1) >> 4
	return
}
func (o Ip6) SetFlow(b []byte, vn uint32) error {
	v, err := ReadB32(b, 1)
	if err != nil {
		return err
	}
	err = WriteB32(b, 1, v|(vn<<4)&((1<<(Ip6Flow+4))-1))
	return err
}
func (o Ip6) Payload(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 4); err != nil {
		return
	}
	v &= ((1 << (Ip6Payload + 0)) - 1) >> 0
	return
}
func (o Ip6) SetPayload(b []byte, vn uint16) error {
	v, err := ReadB16(b, 4)
	if err != nil {
		return err
	}
	err = WriteB16(b, 4, v|(vn<<0)&((1<<(Ip6Payload+0))-1))
	return err
}
func (o Ip6) Next(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 6); err != nil {
		return
	}
	v &= ((1 << (Ip6Next + 0)) - 1) >> 0
	return
}
func (o Ip6) SetNext(b []byte, vn uint8) error {
	v, err := ReadB8(b, 6)
	if err != nil {
		return err
	}
	err = WriteB8(b, 6, v|(vn<<0)&((1<<(Ip6Next+0))-1))
	return err
}
func (o Ip6) HopLimit(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 7); err != nil {
		return
	}
	v &= ((1 << (Ip6HopLimit + 0)) - 1) >> 0
	return
}
func (o Ip6) SetHopLimit(b []byte, vn uint8) error {
	v, err := ReadB8(b, 7)
	if err != nil {
		return err
	}
	err = WriteB8(b, 7, v|(vn<<0)&((1<<(Ip6HopLimit+0))-1))
	return err
}
func (o Ip6) Source1(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 8); err != nil {
		return
	}
	v &= ((1 << (Ip6Source1 + 0)) - 1) >> 0
	return
}
func (o Ip6) SetSource1(b []byte, vn uint64) error {
	v, err := ReadB64(b, 8)
	if err != nil {
		return err
	}
	err = WriteB64(b, 8, v|(vn<<0)&((1<<(Ip6Source1+0))-1))
	return err
}
func (o Ip6) Source2(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 16); err != nil {
		return
	}
	v &= ((1 << (Ip6Source2 + 0)) - 1) >> 0
	return
}
func (o Ip6) SetSource2(b []byte, vn uint64) error {
	v, err := ReadB64(b, 16)
	if err != nil {
		return err
	}
	err = WriteB64(b, 16, v|(vn<<0)&((1<<(Ip6Source2+0))-1))
	return err
}
func (o Ip6) Dest1(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 24); err != nil {
		return
	}
	v &= ((1 << (Ip6Dest1 + 0)) - 1) >> 0
	return
}
func (o Ip6) SetDest1(b []byte, vn uint64) error {
	v, err := ReadB64(b, 24)
	if err != nil {
		return err
	}
	err = WriteB64(b, 24, v|(vn<<0)&((1<<(Ip6Dest1+0))-1))
	return err
}
func (o Ip6) Dest2(b []byte) (v uint64, err error) {
	if v, err = ReadB64(b, 32); err != nil {
		return
	}
	v &= ((1 << (Ip6Dest2 + 0)) - 1) >> 0
	return
}
func (o Ip6) SetDest2(b []byte, vn uint64) error {
	v, err := ReadB64(b, 32)
	if err != nil {
		return err
	}
	err = WriteB64(b, 32, v|(vn<<0)&((1<<(Ip6Dest2+0))-1))
	return err
}
