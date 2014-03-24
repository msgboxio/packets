package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10FlowMod = 64

type Ofp10FlowMod struct {
	// match ofp10_match
	// cookie uint64
	// command uint16
	// idle_timeout uint16
	// hard_timeout uint16
	// priority uint16
	// buffer_id uint32
	// out_port uint16
	// flags uint16
}

func (o Ofp10FlowMod) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 40)
}
func (o Ofp10FlowMod) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 40, vn)
}
func (o Ofp10FlowMod) Command(b []byte) (uint16, error) {
	return packets.ReadB16(b, 48)
}
func (o Ofp10FlowMod) SetCommand(b []byte, vn uint16) error {
	return packets.WriteB16(b, 48, vn)
}
func (o Ofp10FlowMod) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 50)
}
func (o Ofp10FlowMod) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 50, vn)
}
func (o Ofp10FlowMod) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 52)
}
func (o Ofp10FlowMod) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 52, vn)
}
func (o Ofp10FlowMod) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 54)
}
func (o Ofp10FlowMod) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 54, vn)
}
func (o Ofp10FlowMod) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 56)
}
func (o Ofp10FlowMod) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 56, vn)
}
func (o Ofp10FlowMod) OutPort(b []byte) (uint16, error) {
	return packets.ReadB16(b, 60)
}
func (o Ofp10FlowMod) SetOutPort(b []byte, vn uint16) error {
	return packets.WriteB16(b, 60, vn)
}
func (o Ofp10FlowMod) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 62)
}
func (o Ofp10FlowMod) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 62, vn)
}
