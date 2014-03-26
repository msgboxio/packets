package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMatch = 8

type OfpMatch struct {
	// msg_type uint16
	// length uint16
	// oxm_fields [0]byte
	// pad [4]byte
}

func (o OfpMatch) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpMatch) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpMatch) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpMatch) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpMatch) OxmFields(b []byte) ([]byte, error) {
	return b[4:4], nil
}
func (o OfpMatch) SetOxmFields(b []byte, vn []byte) error {
	copy(b[4:4], vn)
	return nil
}
