package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpGroupFeatures = 40

type OfpGroupFeatures struct {
	// types uint32
	// capabilities uint32
	// max_groups [4]uint32
	// actions [4]uint32
}

func (o OfpGroupFeatures) Types(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o OfpGroupFeatures) SetTypes(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o OfpGroupFeatures) Capabilities(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpGroupFeatures) SetCapabilities(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpGroupFeatures) MaxGroups(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 4)
	for i := 0; i < 4; i++ {
		ret[i], err = packets.ReadB32(b, 8+i*4)
	}
	return
}
func (o OfpGroupFeatures) SetMaxGroups(b []byte, vn []uint32) (err error) {
	for i := 0; i < 4; i++ {
		err = packets.WriteB32(b, 8+i*4, vn[i])
	}
	return
}
func (o OfpGroupFeatures) Actions(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 4)
	for i := 0; i < 4; i++ {
		ret[i], err = packets.ReadB32(b, 24+i*4)
	}
	return
}
func (o OfpGroupFeatures) SetActions(b []byte, vn []uint32) (err error) {
	for i := 0; i < 4; i++ {
		err = packets.WriteB32(b, 24+i*4, vn[i])
	}
	return
}
