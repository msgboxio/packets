package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableFeaturePropHeader = 4

type Ofp13TableFeaturePropHeader struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp13TableFeaturePropHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13TableFeaturePropHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13TableFeaturePropHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13TableFeaturePropHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
