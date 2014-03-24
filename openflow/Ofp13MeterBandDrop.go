package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterBandDrop = 16

type Ofp13MeterBandDrop struct {
	// msg_type uint16
	// len uint16
	// rate uint32
	// burst_size uint32
	// pad [4]byte
}

func (o Ofp13MeterBandDrop) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp13MeterBandDrop) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp13MeterBandDrop) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp13MeterBandDrop) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp13MeterBandDrop) Rate(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp13MeterBandDrop) SetRate(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp13MeterBandDrop) BurstSize(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp13MeterBandDrop) SetBurstSize(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
