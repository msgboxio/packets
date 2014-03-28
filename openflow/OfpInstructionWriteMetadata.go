package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpInstructionWriteMetadata = 24

type OfpInstructionWriteMetadata struct {
	// msg_type uint16
	// len uint16
	// pad [4]byte
	// metadata uint64
	// metadata_mask uint64
}

func (o OfpInstructionWriteMetadata) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpInstructionWriteMetadata) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpInstructionWriteMetadata) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpInstructionWriteMetadata) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpInstructionWriteMetadata) Metadata(b []byte) (uint64, error) {
	return packets.ReadB64(b, 8)
}
func (o OfpInstructionWriteMetadata) SetMetadata(b []byte, vn uint64) error {
	return packets.WriteB64(b, 8, vn)
}
func (o OfpInstructionWriteMetadata) MetadataMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 16)
}
func (o OfpInstructionWriteMetadata) SetMetadataMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 16, vn)
}
