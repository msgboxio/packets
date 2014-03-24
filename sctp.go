package packets

const (
	SctpSrc  = 16
	SctpDst  = 16
	SctpVtag = 32
	SctpCsum = 32
)

type Sctp struct {
}

func (o Sctp) Src(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 0); err != nil {
		return
	}
	v &= ((1 << (SctpSrc + 0)) - 1) >> 0
	return
}
func (o Sctp) SetSrc(b []byte, vn uint16) error {
	v, err := ReadB16(b, 0)
	if err != nil {
		return err
	}
	err = WriteB16(b, 0, v|(vn<<0)&((1<<(SctpSrc+0))-1))
	return err
}
func (o Sctp) Dst(b []byte) (v uint16, err error) {
	if v, err = ReadB16(b, 2); err != nil {
		return
	}
	v &= ((1 << (SctpDst + 0)) - 1) >> 0
	return
}
func (o Sctp) SetDst(b []byte, vn uint16) error {
	v, err := ReadB16(b, 2)
	if err != nil {
		return err
	}
	err = WriteB16(b, 2, v|(vn<<0)&((1<<(SctpDst+0))-1))
	return err
}
func (o Sctp) Vtag(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 4); err != nil {
		return
	}
	v &= ((1 << (SctpVtag + 0)) - 1) >> 0
	return
}
func (o Sctp) SetVtag(b []byte, vn uint32) error {
	v, err := ReadB32(b, 4)
	if err != nil {
		return err
	}
	err = WriteB32(b, 4, v|(vn<<0)&((1<<(SctpVtag+0))-1))
	return err
}
func (o Sctp) Csum(b []byte) (v uint32, err error) {
	if v, err = ReadB32(b, 8); err != nil {
		return
	}
	v &= ((1 << (SctpCsum + 0)) - 1) >> 0
	return
}
func (o Sctp) SetCsum(b []byte, vn uint32) error {
	v, err := ReadB32(b, 8)
	if err != nil {
		return err
	}
	err = WriteB32(b, 8, v|(vn<<0)&((1<<(SctpCsum+0))-1))
	return err
}
