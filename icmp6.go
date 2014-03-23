package packets

const (
	icmp6_type     = 8
	icmp6_code     = 8
	icmp6_checksum = 16
	icmp6_specific = 32
)

type icmp6 struct {
}

func (o icmp6) Type(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (icmp6_type + 0)) - 1) >> 0
	return
}
func (o icmp6) SetType(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(icmp6_type+0))-1))
	return err
}
func (o icmp6) Code(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (icmp6_code + 0)) - 1) >> 0
	return
}
func (o icmp6) SetCode(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(icmp6_code+0))-1))
	return err
}
func (o icmp6) Checksum(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 2); err != nil {
		return
	}
	v &= ((1 << (icmp6_checksum + 0)) - 1) >> 0
	return
}
func (o icmp6) SetChecksum(b []byte, vn uint16) error {
	v, err := read_b16(b, 2)
	if err != nil {
		return err
	}
	err = write_b16(b, 2, v|(vn<<0)&((1<<(icmp6_checksum+0))-1))
	return err
}
func (o icmp6) Specific(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 4); err != nil {
		return
	}
	v &= ((1 << (icmp6_specific + 0)) - 1) >> 0
	return
}
func (o icmp6) SetSpecific(b []byte, vn uint32) error {
	v, err := read_b32(b, 4)
	if err != nil {
		return err
	}
	err = write_b32(b, 4, v|(vn<<0)&((1<<(icmp6_specific+0))-1))
	return err
}
