package openflow

import (
	"msgbox.io/packets"
)

const SizeofOfpAsyncConfig = 32

type OfpAsyncConfig struct {
	// header ofp_header
	// packet_in_mask [2]uint32
	// port_status_mask [2]uint32
	// flow_removed_mask [2]uint32
}

func (o OfpAsyncConfig) PacketInMask(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 2)
	for i := 0; i < 2; i++ {
		ret[i], err = packets.ReadB32(b, 8+i*4)
	}
	return
}
func (o OfpAsyncConfig) SetPacketInMask(b []byte, vn []uint32) (err error) {
	for i := 0; i < 2; i++ {
		err = packets.WriteB32(b, 8+i*4, vn[i])
	}
	return
}
func (o OfpAsyncConfig) PortStatusMask(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 2)
	for i := 0; i < 2; i++ {
		ret[i], err = packets.ReadB32(b, 16+i*4)
	}
	return
}
func (o OfpAsyncConfig) SetPortStatusMask(b []byte, vn []uint32) (err error) {
	for i := 0; i < 2; i++ {
		err = packets.WriteB32(b, 16+i*4, vn[i])
	}
	return
}
func (o OfpAsyncConfig) FlowRemovedMask(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 2)
	for i := 0; i < 2; i++ {
		ret[i], err = packets.ReadB32(b, 24+i*4)
	}
	return
}
func (o OfpAsyncConfig) SetFlowRemovedMask(b []byte, vn []uint32) (err error) {
	for i := 0; i < 2; i++ {
		err = packets.WriteB32(b, 24+i*4, vn[i])
	}
	return
}
