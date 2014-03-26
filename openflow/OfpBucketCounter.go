package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpBucketCounter = 16

type OfpBucketCounter struct {
	// packet_count uint64
	// byte_count uint64
}

func (o OfpBucketCounter) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o OfpBucketCounter) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o OfpBucketCounter) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpBucketCounter) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
