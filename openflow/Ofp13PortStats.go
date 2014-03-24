package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13PortStats = 112

type Ofp13PortStats struct {
	// ps ofp11_port_stats
	// duration_sec uint32
	// duration_nsec uint32
}

func (o Ofp13PortStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 104)
}
func (o Ofp13PortStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 104, vn)
}
func (o Ofp13PortStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 108)
}
func (o Ofp13PortStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 108, vn)
}
