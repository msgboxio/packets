package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp13AsyncConfig = 24

type Ofp13AsyncConfig struct {
	// packet_in_mask [2]uint32
	// port_status_mask [2]uint32
	// flow_removed_mask [2]uint32
}

func (o Ofp13AsyncConfig) PacketInMask(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 2)
	for i := 0; i < 2; i++ {
		ret[i], err = packets.ReadB32(b, 0+i*4)
	}
	return
}
func (o Ofp13AsyncConfig) SetPacketInMask(b []byte, vn []uint32) (err error) {
	for i := 0; i < 2; i++ {
		err = packets.WriteB32(b, 0+i*4, vn[i])
	}
	return
}
func (o Ofp13AsyncConfig) PortStatusMask(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 2)
	for i := 0; i < 2; i++ {
		ret[i], err = packets.ReadB32(b, 8+i*4)
	}
	return
}
func (o Ofp13AsyncConfig) SetPortStatusMask(b []byte, vn []uint32) (err error) {
	for i := 0; i < 2; i++ {
		err = packets.WriteB32(b, 8+i*4, vn[i])
	}
	return
}
func (o Ofp13AsyncConfig) FlowRemovedMask(b []byte) (ret []uint32, err error) {
	ret = make([]uint32, 2)
	for i := 0; i < 2; i++ {
		ret[i], err = packets.ReadB32(b, 16+i*4)
	}
	return
}
func (o Ofp13AsyncConfig) SetFlowRemovedMask(b []byte, vn []uint32) (err error) {
	for i := 0; i < 2; i++ {
		err = packets.WriteB32(b, 16+i*4, vn[i])
	}
	return
}
