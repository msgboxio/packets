package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionSetField = 8

type OfpActionSetField struct {
	// msg_type uint16
	// len uint16
	// field [4]byte
}

func (o OfpActionSetField) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionSetField) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionSetField) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionSetField) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionSetField) Field(b []byte) ([]byte, error) {
	return b[4:8], nil
}
func (o OfpActionSetField) SetField(b []byte, vn []byte) error {
	copy(b[4:8], vn)
	return nil
}
