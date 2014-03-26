package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpSwitchConfig = 12

type OfpSwitchConfig struct {
	// header ofp_header
	// flags uint16
	// miss_send_len uint16
}

func (o OfpSwitchConfig) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 8)
}
func (o OfpSwitchConfig) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 8, vn)
}
func (o OfpSwitchConfig) MissSendLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 10)
}
func (o OfpSwitchConfig) SetMissSendLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 10, vn)
}
