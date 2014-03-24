package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterBandDscpRemark = 16

type Ofp13MeterBandDscpRemark struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
	// prec_level uint8
	// pad [3]byte
}

func (o Ofp13MeterBandDscpRemark) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13MeterBandDscpRemark) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13MeterBandDscpRemark) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13MeterBandDscpRemark) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13MeterBandDscpRemark) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterBandDscpRemark) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp13MeterBandDscpRemark) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp13MeterBandDscpRemark) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o Ofp13MeterBandDscpRemark) PrecLevel(b []byte) (uint8, error) {
	return packets.ReadB8(b, 12)
}
func (o Ofp13MeterBandDscpRemark) SetPrecLevel(b []byte, vn uint8) error {
	return packets.WriteB8(b, 12, vn)
}
