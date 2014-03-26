package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMultipartReply = 16

type OfpMultipartReply struct {
	// header ofp_header
	// msg_type uint16
	// flags uint16
	// pad [4]byte
	// body [0]byte
}

func (o OfpMultipartReply) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpMultipartReply) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpMultipartReply) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o OfpMultipartReply) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
func (o OfpMultipartReply) Body(b []byte) ([]byte, error) {
	return b[16:16], nil
}
func (o OfpMultipartReply) SetBody(b []byte, vn []byte) error {
	copy(b[16:16], vn)
	return nil
}
