package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpHeader = 8

type OfpHeader struct {
	// version uint8
	// msg_type uint8
	// length uint16
	// xid uint32
}

func (o OfpHeader) Version(b []byte) (uint8, error) {
	return packets.ReadB8(b, 0)
}
func (o OfpHeader) SetVersion(b []byte, vn uint8) error {
	return packets.WriteB8(b, 0, vn)
}
func (o OfpHeader) MsgType(b []byte) (uint8, error) {
	return packets.ReadB8(b, 1)
}
func (o OfpHeader) SetMsgType(b []byte, vn uint8) error {
	return packets.WriteB8(b, 1, vn)
}
func (o OfpHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpHeader) Xid(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpHeader) SetXid(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
