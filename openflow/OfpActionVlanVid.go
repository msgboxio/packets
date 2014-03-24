package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionVlanVid = 8

type OfpActionVlanVid struct {
	// msg_type uint16
	// len uint16
	// vlan_vid uint16
	// pad [2]byte
}

func (o OfpActionVlanVid) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionVlanVid) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionVlanVid) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionVlanVid) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionVlanVid) VlanVid(b []byte) (uint16, error) {
	return packets.ReadB16(b, 4)
}
func (o OfpActionVlanVid) SetVlanVid(b []byte, vn uint16) error {
	return packets.WriteB16(b, 4, vn)
}
