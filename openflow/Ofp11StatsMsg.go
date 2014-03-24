package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11StatsMsg = 16

type Ofp11StatsMsg struct {
	// header ofp_header
	// msg_type uint16
	// flags uint16
	// pad [4]byte
}

func (o Ofp11StatsMsg) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp11StatsMsg) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o Ofp11StatsMsg) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o Ofp11StatsMsg) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
