package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpTableFeaturePropExperimenter = 12

type OfpTableFeaturePropExperimenter struct {
	// msg_type uint16
	// length uint16
	// experimenter uint32
	// exp_type uint32
	// experimenter_data [0]uint32
}

func (o OfpTableFeaturePropExperimenter) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpTableFeaturePropExperimenter) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpTableFeaturePropExperimenter) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpTableFeaturePropExperimenter) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpTableFeaturePropExperimenter) Experimenter(b []byte) (uint32, error) {
	return packets.ReadB32(b, 4)
}
func (o OfpTableFeaturePropExperimenter) SetExperimenter(b []byte, vn uint32) error {
	return packets.WriteB32(b, 4, vn)
}
func (o OfpTableFeaturePropExperimenter) ExpType(b []byte) (uint32, error) {
	return packets.ReadB32(b, 8)
}
func (o OfpTableFeaturePropExperimenter) SetExpType(b []byte, vn uint32) error {
	return packets.WriteB32(b, 8, vn)
}
func (o OfpTableFeaturePropExperimenter) ExperimenterData(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 0)
	for i := 0; i < 0; i++ {
		ret[i], err = packets.ReadB32(b, 12+i*4)
	}
	return
}
func (o OfpTableFeaturePropExperimenter) SetExperimenterData(b []byte, vn []uint32) (err error) {
	for i := 0; i < 0; i++ {
		err = packets.WriteB32(b, 12+i*4, vn[i])
	}
	return
}
