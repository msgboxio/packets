package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpErrorMsg = 4

type OfpErrorMsg struct {
	// msg_type uint16
	// code uint16
	// data [0]byte
}

func (o OfpErrorMsg) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpErrorMsg) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpErrorMsg) Code(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpErrorMsg) SetCode(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpErrorMsg) Data(b []byte) ([]byte, error) {
	return b[4:4], nil
}
func (o OfpErrorMsg) SetData(b []byte, vn []byte) error {
	copy(b[4:4], vn)
	return nil
}
