package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpDescStats = 1056

type OfpDescStats struct {
	// mfr_desc char[256]
	// hw_desc char[256]
	// sw_desc char[256]
	// serial_num char[32]
	// dp_desc char[256]
}

func (o OfpDescStats) MfrDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 0, 256, '\000')
	return string(slice), err
}
func (o OfpDescStats) SetMfrDesc(b []byte, vn string) error {
	copy(b[0:], vn)
	return nil
}
func (o OfpDescStats) HwDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 1, 256, '\000')
	return string(slice), err
}
func (o OfpDescStats) SetHwDesc(b []byte, vn string) error {
	copy(b[1:], vn)
	return nil
}
func (o OfpDescStats) SwDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 2, 256, '\000')
	return string(slice), err
}
func (o OfpDescStats) SetSwDesc(b []byte, vn string) error {
	copy(b[2:], vn)
	return nil
}
func (o OfpDescStats) SerialNum(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 3, 32, '\000')
	return string(slice), err
}
func (o OfpDescStats) SetSerialNum(b []byte, vn string) error {
	copy(b[3:], vn)
	return nil
}
func (o OfpDescStats) DpDesc(b []byte) (string, error) {
	slice, err := packets.ReadSlice(b, 4, 256, '\000')
	return string(slice), err
}
func (o OfpDescStats) SetDpDesc(b []byte, vn string) error {
	copy(b[4:], vn)
	return nil
}
