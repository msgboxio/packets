package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableFeaturePropOxm = 4

type Ofp13TableFeaturePropOxm struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp13TableFeaturePropOxm) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13TableFeaturePropOxm) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13TableFeaturePropOxm) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13TableFeaturePropOxm) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
