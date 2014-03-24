package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp14RoleStatus = 16

type Ofp14RoleStatus struct {
	// role uint32
	// reason uint8
	// pad [3]byte
	// generation_id uint64
}

func (o Ofp14RoleStatus) Role(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp14RoleStatus) SetRole(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp14RoleStatus) Reason(b []byte) (uint8, error) {
	return packets.ReadB8(b, 4)
}
func (o Ofp14RoleStatus) SetReason(b []byte, vn uint8) error {
	return packets.WriteB8(b, 4, vn)
}
func (o Ofp14RoleStatus) GenerationId(b []byte) (uint64, error) {
	return packets.ReadB64(b, 6)
}
func (o Ofp14RoleStatus) SetGenerationId(b []byte, vn uint64) error {
	return packets.WriteB64(b, 6, vn)
}
