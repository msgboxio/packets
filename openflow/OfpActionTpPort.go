package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionTpPort = 8

type OfpActionTpPort struct {
	// msg_type uint16
	// len uint16
	// tp_port uint16
	// pad [2]byte
}

func (o OfpActionTpPort) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionTpPort) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionTpPort) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionTpPort) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionTpPort) TpPort(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o OfpActionTpPort) SetTpPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
