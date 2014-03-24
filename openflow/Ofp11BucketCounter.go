package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11BucketCounter = 16

type Ofp11BucketCounter struct {
	// packet_count uint64
	// byte_count uint64
}

func (o Ofp11BucketCounter) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o Ofp11BucketCounter) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o Ofp11BucketCounter) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp11BucketCounter) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
