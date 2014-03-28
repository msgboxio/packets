package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpTableFeaturePropActions = 4

type OfpTableFeaturePropActions struct {
	// msg_type uint16
	// length uint16
}

func (o OfpTableFeaturePropActions) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpTableFeaturePropActions) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpTableFeaturePropActions) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpTableFeaturePropActions) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
