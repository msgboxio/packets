package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpMultipartRequest = 16

type OfpMultipartRequest struct {
	// header ofp_header
	// msg_type uint16
	// flags uint16
	// pad [4]byte
	// body [0]byte
}

func (o OfpMultipartRequest) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpMultipartRequest) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpMultipartRequest) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o OfpMultipartRequest) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
func (o OfpMultipartRequest) Body(b []byte) ([]byte, error) {
	return b[16:16], nil
}
func (o OfpMultipartRequest) SetBody(b []byte, vn []byte) error {
	copy(b[16:16], vn)
	return nil
}
