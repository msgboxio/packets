package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11ActionOutput = 16

type Ofp11ActionOutput struct {
	// msg_type uint16
	// len uint16
	// port uint32
	// max_len uint16
	// pad [6]byte
}

func (o Ofp11ActionOutput) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11ActionOutput) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11ActionOutput) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11ActionOutput) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11ActionOutput) Port(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11ActionOutput) SetPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11ActionOutput) MaxLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp11ActionOutput) SetMaxLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
