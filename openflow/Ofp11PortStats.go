package openflow

import (
	"msgbox.io/packets"
)

const sizeofOfp11PortStats = 104

type Ofp11PortStats struct {
	// port_no uint32
	// pad [4]byte
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

func (o Ofp11PortStats) PortNo(b []byte) (uint32, error) {
	return packets.ReadB32(b, 0)
}
func (o Ofp11PortStats) SetPortNo(b []byte, vn uint32) error {
	return packets.WriteB32(b, 0, vn)
}
func (o Ofp11PortStats) RxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 5)
}
func (o Ofp11PortStats) SetRxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 5, vn)
}
func (o Ofp11PortStats) TxPackets(b []byte) (uint64, error) {
	return packets.ReadB64(b, 13)
}
func (o Ofp11PortStats) SetTxPackets(b []byte, vn uint64) error {
	return packets.WriteB64(b, 13, vn)
}
func (o Ofp11PortStats) RxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 21)
}
func (o Ofp11PortStats) SetRxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 21, vn)
}
func (o Ofp11PortStats) TxBytes(b []byte) (uint64, error) {
	return packets.ReadB64(b, 29)
}
func (o Ofp11PortStats) SetTxBytes(b []byte, vn uint64) error {
	return packets.WriteB64(b, 29, vn)
}
func (o Ofp11PortStats) RxDropped(b []byte) (uint64, error) {
	return packets.ReadB64(b, 37)
}
func (o Ofp11PortStats) SetRxDropped(b []byte, vn uint64) error {
	return packets.WriteB64(b, 37, vn)
}
func (o Ofp11PortStats) TxDropped(b []byte) (uint64, error) {
	return packets.ReadB64(b, 45)
}
func (o Ofp11PortStats) SetTxDropped(b []byte, vn uint64) error {
	return packets.WriteB64(b, 45, vn)
}
func (o Ofp11PortStats) RxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 53)
}
func (o Ofp11PortStats) SetRxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 53, vn)
}
func (o Ofp11PortStats) TxErrors(b []byte) (uint64, error) {
	return packets.ReadB64(b, 61)
}
func (o Ofp11PortStats) SetTxErrors(b []byte, vn uint64) error {
	return packets.WriteB64(b, 61, vn)
}
func (o Ofp11PortStats) RxFrameErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 69)
}
func (o Ofp11PortStats) SetRxFrameErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 69, vn)
}
func (o Ofp11PortStats) RxOverErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 77)
}
func (o Ofp11PortStats) SetRxOverErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 77, vn)
}
func (o Ofp11PortStats) RxCrcErr(b []byte) (uint64, error) {
	return packets.ReadB64(b, 85)
}
func (o Ofp11PortStats) SetRxCrcErr(b []byte, vn uint64) error {
	return packets.WriteB64(b, 85, vn)
}
func (o Ofp11PortStats) Collisions(b []byte) (uint64, error) {
	return packets.ReadB64(b, 93)
}
func (o Ofp11PortStats) SetCollisions(b []byte, vn uint64) error {
	return packets.WriteB64(b, 93, vn)
}
