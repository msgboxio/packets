package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpActionVlanPcp = 8

type OfpActionVlanPcp struct {
	// msg_type uint16
	// len uint16
	// vlan_pcp uint8
	// pad [3]byte
}

func (o OfpActionVlanPcp) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionVlanPcp) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionVlanPcp) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionVlanPcp) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionVlanPcp) VlanPcp(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o OfpActionVlanPcp) SetVlanPcp(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
