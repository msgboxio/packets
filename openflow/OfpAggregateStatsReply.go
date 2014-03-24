package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpAggregateStatsReply = 24

type OfpAggregateStatsReply struct {
	// packet_count uint64
	// byte_count uint64
	// flow_count uint32
	// pad [4]byte
}

func (o OfpAggregateStatsReply) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o OfpAggregateStatsReply) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o OfpAggregateStatsReply) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpAggregateStatsReply) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpAggregateStatsReply) FlowCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o OfpAggregateStatsReply) SetFlowCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
