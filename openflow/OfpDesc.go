package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpDesc = 1056

type OfpDesc struct {
	// mfr_desc char[256]
	// hw_desc char[256]
	// sw_desc char[256]
	// serial_num char[32]
	// dp_desc char[256]
}

func (o OfpDesc) MfrDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 0, 256, '\000')
	return string(slice), err
}
func (o OfpDesc) SetMfrDesc(b []byte, vn string) error {
	copy(b[0:], vn)
	return nil
}
func (o OfpDesc) HwDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 256, 256, '\000')
	return string(slice), err
}
func (o OfpDesc) SetHwDesc(b []byte, vn string) error {
	copy(b[256:], vn)
	return nil
}
func (o OfpDesc) SwDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 512, 256, '\000')
	return string(slice), err
}
func (o OfpDesc) SetSwDesc(b []byte, vn string) error {
	copy(b[512:], vn)
	return nil
}
func (o OfpDesc) SerialNum(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 768, 32, '\000')
	return string(slice), err
}
func (o OfpDesc) SetSerialNum(b []byte, vn string) error {
	copy(b[768:], vn)
	return nil
}
func (o OfpDesc) DpDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 800, 256, '\000')
	return string(slice), err
}
func (o OfpDesc) SetDpDesc(b []byte, vn string) error {
	copy(b[800:], vn)
	return nil
}
