package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpBucket = 16

type OfpBucket struct {
	// len uint16
	// weight uint16
	// watch_port uint32
	// watch_group uint32
	// pad [4]byte
}

func (o OfpBucket) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpBucket) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpBucket) Weight(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpBucket) SetWeight(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpBucket) WatchPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpBucket) SetWatchPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpBucket) WatchGroup(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpBucket) SetWatchGroup(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
