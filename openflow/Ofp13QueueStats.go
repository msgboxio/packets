package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13QueueStats = 40

type Ofp13QueueStats struct {
	// qs ofp11_queue_stats
	// duration_sec uint32
	// duration_nsec uint32
}

func (o Ofp13QueueStats) DurationSec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o Ofp13QueueStats) SetDurationSec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o Ofp13QueueStats) DurationNsec(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o Ofp13QueueStats) SetDurationNsec(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
