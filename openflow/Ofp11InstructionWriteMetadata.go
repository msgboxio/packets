package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11InstructionWriteMetadata = 24

type Ofp11InstructionWriteMetadata struct {
	// msg_type uint16
	// len uint16
	// pad [4]byte
	// metadata uint64
	// metadata_mask uint64
}

func (o Ofp11InstructionWriteMetadata) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp11InstructionWriteMetadata) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp11InstructionWriteMetadata) Len(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o Ofp11InstructionWriteMetadata) SetLen(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o Ofp11InstructionWriteMetadata) Metadata(b []byte) (uint64, error) {
	return packets.ReadB64(b, 5)
}
func (o Ofp11InstructionWriteMetadata) SetMetadata(b []byte, vn uint64) error {
	return packets.WriteB64(b, 5, vn)
}
func (o Ofp11InstructionWriteMetadata) MetadataMask(b []byte) (uint64, error) {
	return packets.ReadB64(b, 13)
}
func (o Ofp11InstructionWriteMetadata) SetMetadataMask(b []byte, vn uint64) error {
	return packets.WriteB64(b, 13, vn)
}
