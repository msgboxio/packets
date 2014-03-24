package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp14RolePropHeader = 4

type Ofp14RolePropHeader struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp14RolePropHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp14RolePropHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp14RolePropHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp14RolePropHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
