package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13PacketIn = 16

type Ofp13PacketIn struct {
	// pi ofp12_packet_in
	// cookie uint64
}

func (o Ofp13PacketIn) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp13PacketIn) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
