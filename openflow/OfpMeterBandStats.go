package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpMeterBandStats = 16

type OfpMeterBandStats struct {
	// packet_band_count uint64
	// byte_band_count uint64
}

func (o OfpMeterBandStats) PacketBandCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o OfpMeterBandStats) SetPacketBandCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o OfpMeterBandStats) ByteBandCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpMeterBandStats) SetByteBandCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
