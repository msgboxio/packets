package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpSwitchFeatures = 32

type OfpSwitchFeatures struct {
	// header ofp_header
	// datapath_id uint64
	// n_buffers uint32
	// n_tables uint8
	// auxiliary_id uint8
	// pad [2]byte
	// capabilities uint32
	// reserved uint32
}

func (o OfpSwitchFeatures) DatapathId(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpSwitchFeatures) SetDatapathId(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpSwitchFeatures) NBuffers(b []byte) (uint32, error) {
	return packets.ReadB32(b, 16)
}
func (o OfpSwitchFeatures) SetNBuffers(b []byte, vn uint32) error {
	return packets.WriteB32(b, 16, vn)
}
func (o OfpSwitchFeatures) NTables(b []byte) (uint8, error) {
	return packets.ReadB8(b, 20)
}
func (o OfpSwitchFeatures) SetNTables(b []byte, vn uint8) error {
	return packets.WriteB8(b, 20, vn)
}
func (o OfpSwitchFeatures) AuxiliaryId(b []byte) (uint8, error) {
	return packets.ReadB8(b, 21)
}
func (o OfpSwitchFeatures) SetAuxiliaryId(b []byte, vn uint8) error {
	return packets.WriteB8(b, 21, vn)
}
func (o OfpSwitchFeatures) Capabilities(b []byte) (uint32, error) {
	return packets.ReadB32(b, 24)
}
func (o OfpSwitchFeatures) SetCapabilities(b []byte, vn uint32) error {
	return packets.WriteB32(b, 24, vn)
}
func (o OfpSwitchFeatures) Reserved(b []byte) (uint32, error) {
	return packets.ReadB32(b, 28)
}
func (o OfpSwitchFeatures) SetReserved(b []byte, vn uint32) error {
	return packets.WriteB32(b, 28, vn)
}
