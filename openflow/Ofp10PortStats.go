package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp10PortStats = 104

type Ofp10PortStats struct {
	// port_no uint16
	// pad [6]byte
	// rx_packets uint64
	// tx_packets uint64
	// rx_bytes uint64
	// tx_bytes uint64
	// rx_dropped uint64
	// tx_dropped uint64
	// rx_errors uint64
	// tx_errors uint64
	// rx_frame_err uint64
	// rx_over_err uint64
	// rx_crc_err uint64
	// collisions uint64
}

func (o Ofp10PortStats) PortNo(b []byte) (uint16, error) {
	return packets.ReadB16(b, 0)
}
func (o Ofp10PortStats) SetPortNo(b []byte, vn uint16) error {
	return packets.WriteB16(b, 0, vn)
}
func (o Ofp10PortStats) RxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 3)
}
func (o Ofp10PortStats) SetRxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 3, vn)
}
func (o Ofp10PortStats) TxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 11)
}
func (o Ofp10PortStats) SetTxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 11, vn)
}
func (o Ofp10PortStats) RxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 19)
}
func (o Ofp10PortStats) SetRxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 19, vn)
}
func (o Ofp10PortStats) TxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 27)
}
func (o Ofp10PortStats) SetTxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 27, vn)
}
func (o Ofp10PortStats) RxDropped(b []byte) (uint64, error) {
	return packets.ReadB64(b, 35)
}
func (o Ofp10PortStats) SetRxDropped(b []byte, vn uint64) error {
	return packets.WriteB64(b, 35, vn)
}
func (o Ofp10PortStats) TxDropped(b []byte) (uint64, error) {
	return packets.ReadB64(b, 43)
}
func (o Ofp10PortStats) SetTxDropped(b []byte, vn uint64) error {
	return packets.WriteB64(b, 43, vn)
}
func (o Ofp10PortStats) RxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 51)
}
func (o Ofp10PortStats) SetRxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 51, vn)
}
func (o Ofp10PortStats) TxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 59)
}
func (o Ofp10PortStats) SetTxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 59, vn)
}
func (o Ofp10PortStats) RxFrameErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 67)
}
func (o Ofp10PortStats) SetRxFrameErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 67, vn)
}
func (o Ofp10PortStats) RxOverErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 75)
}
func (o Ofp10PortStats) SetRxOverErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 75, vn)
}
func (o Ofp10PortStats) RxCrcErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 83)
}
func (o Ofp10PortStats) SetRxCrcErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 83, vn)
}
func (o Ofp10PortStats) Collisions(b []byte) (uint64, error) {
	return packets.ReadB64(b, 91)
}
func (o Ofp10PortStats) SetCollisions(b []byte, vn uint64) error {
	return packets.WriteB64(b, 91, vn)
}
