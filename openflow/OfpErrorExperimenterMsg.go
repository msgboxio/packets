package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpErrorExperimenterMsg = 16

type OfpErrorExperimenterMsg struct {
	// header ofp_header
	// msg_type uint16
	// exp_type uint16
	// experimenter uint32
	// data [0]byte
}

func (o OfpErrorExperimenterMsg) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpErrorExperimenterMsg) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpErrorExperimenterMsg) ExpType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o OfpErrorExperimenterMsg) SetExpType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
func (o OfpErrorExperimenterMsg) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpErrorExperimenterMsg) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
func (o OfpErrorExperimenterMsg) Data(b []byte) ([]byte, error) {
	return b[16:16], nil
}
func (o OfpErrorExperimenterMsg) SetData(b []byte, vn []byte) error {
	copy(b[16:16], vn)
	return nil
}
