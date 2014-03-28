package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpFlowMod = 56

type OfpFlowMod struct {
	// header ofp_header
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
	// match ofp_match
}

func (o OfpFlowMod) Cookie(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpFlowMod) SetCookie(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpFlowMod) CookieMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpFlowMod) SetCookieMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
func (o OfpFlowMod) TableId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 24)
}
func (o OfpFlowMod) SetTableId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 24, vn)
}
func (o OfpFlowMod) Command(b []byte) (uint8, error) {
	return packets.ReadB8(b, 25)
}
func (o OfpFlowMod) SetCommand(b []byte, vn uint8) error {
	return packets.WriteB8(b, 25, vn)
}
func (o OfpFlowMod) IdleTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 26)
}
func (o OfpFlowMod) SetIdleTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 26, vn)
}
func (o OfpFlowMod) HardTimeout(b []byte) (uint16, error) {
	return packets.ReadB16(b, 28)
}
func (o OfpFlowMod) SetHardTimeout(b []byte, vn uint16) error {
	return packets.WriteB16(b, 28, vn)
}
func (o OfpFlowMod) Priority(b []byte) (uint16, error) {
	return packets.ReadB16(b, 30)
}
func (o OfpFlowMod) SetPriority(b []byte, vn uint16) error {
	return packets.WriteB16(b, 30, vn)
}
func (o OfpFlowMod) BufferId(b []byte) (uint32, error) {
	return packets.ReadB32(b, 32)
}
func (o OfpFlowMod) SetBufferId(b []byte, vn uint32) error {
	return packets.WriteB32(b, 32, vn)
}
func (o OfpFlowMod) OutPort(b []byte) (uint32, error) {
	return packets.ReadB32(b, 36)
}
func (o OfpFlowMod) SetOutPort(b []byte, vn uint32) error {
	return packets.WriteB32(b, 36, vn)
}
func (o OfpFlowMod) OutGroup(b []byte) (uint32, error) {
	return packets.ReadB32(b, 40)
}
func (o OfpFlowMod) SetOutGroup(b []byte, vn uint32) error {
	return packets.WriteB32(b, 40, vn)
}
func (o OfpFlowMod) Flags(b []byte) (uint16, error) {
	return packets.ReadB16(b, 44)
}
func (o OfpFlowMod) SetFlags(b []byte, vn uint16) error {
	return packets.WriteB16(b, 44, vn)
}
