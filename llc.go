package packets

const (
	LlcDsap = 8
	LlcSsap = 8
	LlcCntl = 8
)

type Llc struct {
}

func (o Llc) Dsap(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (LlcDsap + 0)) - 1) >> 0
	return
}
func (o Llc) SetDsap(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(LlcDsap+0))-1))
	return err
}
func (o Llc) Ssap(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (LlcSsap + 0)) - 1) >> 0
	return
}
func (o Llc) SetSsap(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(LlcSsap+0))-1))
	return err
}
func (o Llc) Cntl(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 2); err != nil {
		return
	}
	v &= ((1 << (LlcCntl + 0)) - 1) >> 0
	return
}
func (o Llc) SetCntl(b []byte, vn uint8) error {
	v, err := ReadB8(b, 2)
	if err != nil {
		return err
	}
	err = WriteB8(b, 2, v|(vn<<0)&((1<<(LlcCntl+0))-1))
	return err
}
