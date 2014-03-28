package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpRoleRequest = 24

type OfpRoleRequest struct {
	// header ofp_header
	// role uint32
	// pad [4]byte
	// generation_id uint64
}

func (o OfpRoleRequest) Role(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpRoleRequest) SetRole(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpRoleRequest) GenerationId(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpRoleRequest) SetGenerationId(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
