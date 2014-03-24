package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionDlAddr = 16

type OfpActionDlAddr struct {
	// msg_type uint16
	// len uint16
	// dl_addr [6]byte
	// pad [6]byte
}

func (o OfpActionDlAddr) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionDlAddr) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionDlAddr) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionDlAddr) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionDlAddr) DlAddr(b []byte) ([]byte, error) {
	return b[4:10], nil
}
func (o OfpActionDlAddr) SetDlAddr(b []byte, vn []byte) error {
	copy(b[4:10], vn)
	return nil
}
