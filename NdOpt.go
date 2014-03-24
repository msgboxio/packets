package packets

const (
	NdOptNdtype = 8
	NdOptLength = 8
)

type NdOpt struct {
}

func (o NdOpt) Ndtype(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (NdOptNdtype + 0)) - 1) >> 0
	return
}
func (o NdOpt) SetNdtype(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(NdOptNdtype+0))-1))
	return err
}
func (o NdOpt) Length(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (NdOptLength + 0)) - 1) >> 0
	return
}
func (o NdOpt) SetLength(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(NdOptLength+0))-1))
	return err
}
