package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMeterBandDscpRemark = 16

type OfpMeterBandDscpRemark struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
	// prec_level uint8
	// pad [3]byte
}

func (o OfpMeterBandDscpRemark) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpMeterBandDscpRemark) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpMeterBandDscpRemark) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpMeterBandDscpRemark) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpMeterBandDscpRemark) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpMeterBandDscpRemark) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpMeterBandDscpRemark) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpMeterBandDscpRemark) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpMeterBandDscpRemark) PrecLevel(b []byte) (uint8, error) {
	return packets.ReadB8(b, 12)
}
func (o OfpMeterBandDscpRemark) SetPrecLevel(b []byte, vn uint8) error {
	return packets.WriteB8(b, 12, vn)
}
