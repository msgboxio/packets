package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10ActionOutput = 8

type Ofp10ActionOutput struct {
	// msg_type uint16
	// len uint16
	// port uint16
	// max_len uint16
}

func (o Ofp10ActionOutput) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10ActionOutput) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10ActionOutput) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp10ActionOutput) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp10ActionOutput) Port(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o Ofp10ActionOutput) SetPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
func (o Ofp10ActionOutput) MaxLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 6)
}
func (o Ofp10ActionOutput) SetMaxLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 6, vn)
}
