package packets

const (
	Ip6ExtNext = 8
	Ip6ExtLen  = 8
)

type Ip6Ext struct {
}

func (o Ip6Ext) Next(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 0); err != nil {
		return
	}
	v &= ((1 << (Ip6ExtNext + 0)) - 1) >> 0
	return
}
func (o Ip6Ext) SetNext(b []byte, vn uint8) error {
	v, err := ReadB8(b, 0)
	if err != nil {
		return err
	}
	err = WriteB8(b, 0, v|(vn<<0)&((1<<(Ip6ExtNext+0))-1))
	return err
}
func (o Ip6Ext) Len(b []byte) (v uint8, err error) {
	if v, err = ReadB8(b, 1); err != nil {
		return
	}
	v &= ((1 << (Ip6ExtLen + 0)) - 1) >> 0
	return
}
func (o Ip6Ext) SetLen(b []byte, vn uint8) error {
	v, err := ReadB8(b, 1)
	if err != nil {
		return err
	}
	err = WriteB8(b, 1, v|(vn<<0)&((1<<(Ip6ExtLen+0))-1))
	return err
}
