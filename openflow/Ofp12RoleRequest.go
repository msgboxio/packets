package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp12RoleRequest = 16

type Ofp12RoleRequest struct {
	// role uint32
	// pad [4]byte
	// generation_id uint64
}

func (o Ofp12RoleRequest) Role(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp12RoleRequest) SetRole(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp12RoleRequest) GenerationId(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o Ofp12RoleRequest) SetGenerationId(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
