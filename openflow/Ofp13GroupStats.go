package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13GroupStats = 40

type Ofp13GroupStats struct {
	// gs ofp11_group_stats
	// duration_sec uint32
	// duration_nsec uint32
}

func (o Ofp13GroupStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o Ofp13GroupStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o Ofp13GroupStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o Ofp13GroupStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
