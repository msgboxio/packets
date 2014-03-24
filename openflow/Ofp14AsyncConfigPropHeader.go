package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp14AsyncConfigPropHeader = 4

type Ofp14AsyncConfigPropHeader struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp14AsyncConfigPropHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp14AsyncConfigPropHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp14AsyncConfigPropHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp14AsyncConfigPropHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
