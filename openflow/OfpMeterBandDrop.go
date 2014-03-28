package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpMeterBandDrop = 16

type OfpMeterBandDrop struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
	// pad [4]byte
}

func (o OfpMeterBandDrop) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpMeterBandDrop) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpMeterBandDrop) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpMeterBandDrop) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpMeterBandDrop) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpMeterBandDrop) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpMeterBandDrop) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpMeterBandDrop) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
