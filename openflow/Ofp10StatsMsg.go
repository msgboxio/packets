package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10StatsMsg = 12

type Ofp10StatsMsg struct {
	// header ofp_header
	// msg_type uint16
	// flags uint16
}

func (o Ofp10StatsMsg) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o Ofp10StatsMsg) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o Ofp10StatsMsg) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o Ofp10StatsMsg) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
