package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterBandExperimenter = 16

type Ofp13MeterBandExperimenter struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
	// experimenter uint32
}

func (o Ofp13MeterBandExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13MeterBandExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13MeterBandExperimenter) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13MeterBandExperimenter) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13MeterBandExperimenter) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterBandExperimenter) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp13MeterBandExperimenter) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp13MeterBandExperimenter) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp13MeterBandExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 12)
}
func (o Ofp13MeterBandExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 12, vn)
}
