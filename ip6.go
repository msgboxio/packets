package packets

const (
	ip6_version   = 4
	ip6_tc1       = 4
	ip6_tc2       = 4
	ip6_flow      = 20
	ip6_payload   = 16
	ip6_next      = 8
	ip6_hop_limit = 8
	ip6_source1   = 64
	ip6_source2   = 64
	ip6_dest1     = 64
	ip6_dest2     = 64
)

type ip6 struct {
}

func (o ip6) Version(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (ip6_version + 0)) - 1) >> 0
	return
}
func (o ip6) SetVersion(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<0)&((1<<(ip6_version+0))-1))
	return err
}
func (o ip6) Tc1(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 0); err != nil {
		return
	}
	v &= ((1 << (ip6_tc1 + 4)) - 1) >> 4
	return
}
func (o ip6) SetTc1(b []byte, vn uint8) error {
	v, err := read_b8(b, 0)
	if err != nil {
		return err
	}
	err = write_b8(b, 0, v|(vn<<4)&((1<<(ip6_tc1+4))-1))
	return err
}
func (o ip6) Tc2(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 1); err != nil {
		return
	}
	v &= ((1 << (ip6_tc2 + 0)) - 1) >> 0
	return
}
func (o ip6) SetTc2(b []byte, vn uint8) error {
	v, err := read_b8(b, 1)
	if err != nil {
		return err
	}
	err = write_b8(b, 1, v|(vn<<0)&((1<<(ip6_tc2+0))-1))
	return err
}
func (o ip6) Flow(b []byte) (v uint32, err error) {
	if v, err = read_b32(b, 1); err != nil {
		return
	}
	v &= ((1 << (ip6_flow + 4)) - 1) >> 4
	return
}
func (o ip6) SetFlow(b []byte, vn uint32) error {
	v, err := read_b32(b, 1)
	if err != nil {
		return err
	}
	err = write_b32(b, 1, v|(vn<<4)&((1<<(ip6_flow+4))-1))
	return err
}
func (o ip6) Payload(b []byte) (v uint16, err error) {
	if v, err = read_b16(b, 4); err != nil {
		return
	}
	v &= ((1 << (ip6_payload + 0)) - 1) >> 0
	return
}
func (o ip6) SetPayload(b []byte, vn uint16) error {
	v, err := read_b16(b, 4)
	if err != nil {
		return err
	}
	err = write_b16(b, 4, v|(vn<<0)&((1<<(ip6_payload+0))-1))
	return err
}
func (o ip6) Next(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 6); err != nil {
		return
	}
	v &= ((1 << (ip6_next + 0)) - 1) >> 0
	return
}
func (o ip6) SetNext(b []byte, vn uint8) error {
	v, err := read_b8(b, 6)
	if err != nil {
		return err
	}
	err = write_b8(b, 6, v|(vn<<0)&((1<<(ip6_next+0))-1))
	return err
}
func (o ip6) Hop_limit(b []byte) (v uint8, err error) {
	if v, err = read_b8(b, 7); err != nil {
		return
	}
	v &= ((1 << (ip6_hop_limit + 0)) - 1) >> 0
	return
}
func (o ip6) SetHop_limit(b []byte, vn uint8) error {
	v, err := read_b8(b, 7)
	if err != nil {
		return err
	}
	err = write_b8(b, 7, v|(vn<<0)&((1<<(ip6_hop_limit+0))-1))
	return err
}
func (o ip6) Source1(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 8); err != nil {
		return
	}
	v &= ((1 << (ip6_source1 + 0)) - 1) >> 0
	return
}
func (o ip6) SetSource1(b []byte, vn uint64) error {
	v, err := read_b64(b, 8)
	if err != nil {
		return err
	}
	err = write_b64(b, 8, v|(vn<<0)&((1<<(ip6_source1+0))-1))
	return err
}
func (o ip6) Source2(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 16); err != nil {
		return
	}
	v &= ((1 << (ip6_source2 + 0)) - 1) >> 0
	return
}
func (o ip6) SetSource2(b []byte, vn uint64) error {
	v, err := read_b64(b, 16)
	if err != nil {
		return err
	}
	err = write_b64(b, 16, v|(vn<<0)&((1<<(ip6_source2+0))-1))
	return err
}
func (o ip6) Dest1(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 24); err != nil {
		return
	}
	v &= ((1 << (ip6_dest1 + 0)) - 1) >> 0
	return
}
func (o ip6) SetDest1(b []byte, vn uint64) error {
	v, err := read_b64(b, 24)
	if err != nil {
		return err
	}
	err = write_b64(b, 24, v|(vn<<0)&((1<<(ip6_dest1+0))-1))
	return err
}
func (o ip6) Dest2(b []byte) (v uint64, err error) {
	if v, err = read_b64(b, 32); err != nil {
		return
	}
	v &= ((1 << (ip6_dest2 + 0)) - 1) >> 0
	return
}
func (o ip6) SetDest2(b []byte, vn uint64) error {
	v, err := read_b64(b, 32)
	if err != nil {
		return err
	}
	err = write_b64(b, 32, v|(vn<<0)&((1<<(ip6_dest2+0))-1))
	return err
}
