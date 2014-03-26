package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionOutput = 16

type OfpActionOutput struct {
	// msg_type uint16
	// len uint16
	// port uint32
	// max_len uint16
	// pad [6]byte
}

func (o OfpActionOutput) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionOutput) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionOutput) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionOutput) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionOutput) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpActionOutput) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpActionOutput) MaxLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpActionOutput) SetMaxLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
