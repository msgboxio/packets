package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11FlowMod = 40

type Ofp11FlowMod struct {
	// cookie uint64
	// cookie_mask uint64
	// table_id uint8
	// command uint8
	// idle_timeout uint16
	// hard_timeout uint16
	// priority uint16
	// buffer_id uint32
	// out_port uint32
	// out_group uint32
	// flags uint16
	// pad [2]byte
}

func (o Ofp11FlowMod) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 0)
}
func (o Ofp11FlowMod) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 0, vn)
}
func (o Ofp11FlowMod) CookieMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp11FlowMod) SetCookieMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o Ofp11FlowMod) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 16)
}
func (o Ofp11FlowMod) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 16, vn)
}
func (o Ofp11FlowMod) Command(b []byte) (uint8, error) {
	return packets.ReadB8(b, 17)
}
func (o Ofp11FlowMod) SetCommand(b []byte, vn uint8) error {
	return packets.WriteB8(b, 17, vn)
}
func (o Ofp11FlowMod) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 18)
}
func (o Ofp11FlowMod) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 18, vn)
}
func (o Ofp11FlowMod) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 20)
}
func (o Ofp11FlowMod) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 20, vn)
}
func (o Ofp11FlowMod) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 22)
}
func (o Ofp11FlowMod) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 22, vn)
}
func (o Ofp11FlowMod) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o Ofp11FlowMod) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
func (o Ofp11FlowMod) OutPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 28)
}
func (o Ofp11FlowMod) SetOutPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 28, vn)
}
func (o Ofp11FlowMod) OutGroup(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o Ofp11FlowMod) SetOutGroup(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o Ofp11FlowMod) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 36)
}
func (o Ofp11FlowMod) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 36, vn)
}
