package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterBandHeader = 12

type Ofp13MeterBandHeader struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
}

func (o Ofp13MeterBandHeader) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13MeterBandHeader) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13MeterBandHeader) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13MeterBandHeader) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13MeterBandHeader) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterBandHeader) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp13MeterBandHeader) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp13MeterBandHeader) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
