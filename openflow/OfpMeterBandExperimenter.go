package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpMeterBandExperimenter = 16

type OfpMeterBandExperimenter struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
	// experimenter uint32
}

func (o OfpMeterBandExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpMeterBandExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpMeterBandExperimenter) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpMeterBandExperimenter) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpMeterBandExperimenter) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpMeterBandExperimenter) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpMeterBandExperimenter) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpMeterBandExperimenter) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpMeterBandExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o OfpMeterBandExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
