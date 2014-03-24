package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13MeterBandStats = 16

type Ofp13MeterBandStats struct {
	// packet_band_count uint64
	// byte_band_count uint64
}

func (o Ofp13MeterBandStats) PacketBandCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o Ofp13MeterBandStats) SetPacketBandCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o Ofp13MeterBandStats) ByteBandCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp13MeterBandStats) SetByteBandCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
