package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableFeaturePropInstructions = 4

type Ofp13TableFeaturePropInstructions struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp13TableFeaturePropInstructions) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13TableFeaturePropInstructions) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13TableFeaturePropInstructions) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13TableFeaturePropInstructions) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
