package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMeterBandHeader = 12

type OfpMeterBandHeader struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
}

func (o OfpMeterBandHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpMeterBandHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpMeterBandHeader) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpMeterBandHeader) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpMeterBandHeader) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpMeterBandHeader) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpMeterBandHeader) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpMeterBandHeader) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
