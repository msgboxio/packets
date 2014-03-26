package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpTableFeaturePropNextTables = 4

type OfpTableFeaturePropNextTables struct {
	// msg_type uint16
	// length uint16
	// next_table_ids [0]byte
}

func (o OfpTableFeaturePropNextTables) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpTableFeaturePropNextTables) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpTableFeaturePropNextTables) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpTableFeaturePropNextTables) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpTableFeaturePropNextTables) NextTableIds(b []byte) ([]byte, error) {
	return b[4:4], nil
}
func (o OfpTableFeaturePropNextTables) SetNextTableIds(b []byte, vn []byte) error {
	copy(b[4:4], vn)
	return nil
}
