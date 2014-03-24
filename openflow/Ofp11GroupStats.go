package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11GroupStats = 32

type Ofp11GroupStats struct {
	// length uint16
	// pad [2]byte
	// group_id uint32
	// ref_count uint32
	// pad2 [4]byte
	// packet_count uint64
	// byte_count uint64
}

func (o Ofp11GroupStats) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11GroupStats) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11GroupStats) GroupId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 3)
}
func (o Ofp11GroupStats) SetGroupId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 3, vn)
}
func (o Ofp11GroupStats) RefCount(b []byte) (uint32, error) {
	return packets.ReadB32(b, 7)
}
func (o Ofp11GroupStats) SetRefCount(b []byte, vn uint32) error {
	return packets.WriteB32(b, 7, vn)
}
func (o Ofp11GroupStats) Pad2(b []byte) ([]byte, error) {
	return b[11:15], nil
}
func (o Ofp11GroupStats) SetPad2(b []byte, vn []byte) error {
	copy(b[11:15], vn)
	return nil
}
func (o Ofp11GroupStats) PacketCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 12)
}
func (o Ofp11GroupStats) SetPacketCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 12, vn)
}
func (o Ofp11GroupStats) ByteCount(b []byte) (uint64, error) {
	return packets.ReadB64(b, 20)
}
func (o Ofp11GroupStats) SetByteCount(b []byte, vn uint64) error {
	return packets.WriteB64(b, 20, vn)
}
