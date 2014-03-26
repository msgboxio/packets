package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpErrorMsg = 12

type OfpErrorMsg struct {
	// header ofp_header
	// msg_type uint16
	// code uint16
	// data [0]byte
}

func (o OfpErrorMsg) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpErrorMsg) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpErrorMsg) Code(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o OfpErrorMsg) SetCode(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
func (o OfpErrorMsg) Data(b []byte) ([]byte, error) {
	return b[12:12], nil
}
func (o OfpErrorMsg) SetData(b []byte, vn []byte) error {
	copy(b[12:12], vn)
	return nil
}
