package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpActionExperimenterHeader = 8

type OfpActionExperimenterHeader struct {
	// msg_type uint16
	// len uint16
	// experimenter uint32
}

func (o OfpActionExperimenterHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpActionExperimenterHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpActionExperimenterHeader) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpActionExperimenterHeader) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpActionExperimenterHeader) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpActionExperimenterHeader) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
