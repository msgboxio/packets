package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11Bucket = 16

type Ofp11Bucket struct {
	// len uint16
	// weight uint16
	// watch_port uint32
	// watch_group uint32
	// pad [4]byte
}

func (o Ofp11Bucket) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11Bucket) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11Bucket) Weight(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11Bucket) SetWeight(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11Bucket) WatchPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o Ofp11Bucket) SetWatchPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o Ofp11Bucket) WatchGroup(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o Ofp11Bucket) SetWatchGroup(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
