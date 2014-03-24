package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableFeaturePropActions = 4

type Ofp13TableFeaturePropActions struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp13TableFeaturePropActions) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13TableFeaturePropActions) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13TableFeaturePropActions) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13TableFeaturePropActions) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
