package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13TableFeaturePropNextTables = 4

type Ofp13TableFeaturePropNextTables struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp13TableFeaturePropNextTables) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13TableFeaturePropNextTables) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13TableFeaturePropNextTables) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13TableFeaturePropNextTables) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
