package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpSwitchConfig = 4

type OfpSwitchConfig struct {
	// flags uint16
	// miss_send_len uint16
}

func (o OfpSwitchConfig) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpSwitchConfig) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpSwitchConfig) MissSendLen(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpSwitchConfig) SetMissSendLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
