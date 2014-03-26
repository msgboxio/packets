package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfpTableFeaturePropOxm = 4

type OfpTableFeaturePropOxm struct {
	// msg_type uint16
	// length uint16
	// oxm_ids [0]uint32
}

func (o OfpTableFeaturePropOxm) MsgType(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o OfpTableFeaturePropOxm) SetMsgType(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o OfpTableFeaturePropOxm) Length(b []byte) (uint16, error) {
	return packets.ReadB16(b, 2)
}
func (o OfpTableFeaturePropOxm) SetLength(b []byte, vn uint16) error {
	return packets.WriteB16(b, 2, vn)
}
func (o OfpTableFeaturePropOxm) OxmIds(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 0)
	for i := 0; i < 0; i++ {
		ret[i], err = packets.ReadB32(b, 4+i*4)
	}
	return
}
func (o OfpTableFeaturePropOxm) SetOxmIds(b []byte, vn []uint32) (err error) {
	for i := 0; i < 0; i++ {
		err = packets.WriteB32(b, 4+i*4, vn[i])
	}
	return
}
