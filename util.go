package packets

import (
	"encoding/binary"
	"io"
)

func read_b64(b []byte, offset int) (uint64, error) {
	if len(b) < 8+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint64(b[offset:8]), nil
}
func read_b32(b []byte, offset int) (uint32, error) {
	if len(b) < 4+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint32(b[offset:4]), nil
}
func read_b16(b []byte, offset int) (uint16, error) {
	if len(b) < 2+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint16(b[offset:2]), nil
}
func read_b8(b []byte, offset int) (uint8, error) {
	if len(b) < 1+offset {
		return 0, io.EOF
	}
	return b[offset], nil
}

func write_b64(b []byte, offset int, v uint64) error {
	if len(b) < 8+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint64(b[offset:8], v)
	return nil
}
func write_b32(b []byte, offset int, v uint32) error {
	if len(b) < 4+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint32(b[offset:4], v)
	return nil
}
func write_b16(b []byte, offset int, v uint16) error {
	if len(b) < 2+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint16(b[offset:2], v)
	return nil
}
func write_b8(b []byte, offset int, v uint8) error {
	if len(b) < offset {
		return io.EOF
	}
	b[offset] = v
	return nil
}
