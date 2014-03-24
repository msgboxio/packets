package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11MatchHeader = 4

type Ofp11MatchHeader struct {
	// msg_type uint16
	// length uint16
}

func (o Ofp11MatchHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11MatchHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11MatchHeader) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11MatchHeader) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
