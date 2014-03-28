package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpHelloElemHeader = 4

type OfpHelloElemHeader struct {
	// msg_type uint16
	// length uint16
}

func (o OfpHelloElemHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpHelloElemHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpHelloElemHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpHelloElemHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
