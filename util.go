package packets

import (
	"bytes"
	"encoding/binary"
	"io"
)

func ReadB64(b []byte, offset int) (uint64, error) {
	if len(b) < 8+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint64(b[offset:8]), nil
}
func ReadB32(b []byte, offset int) (uint32, error) {
	if len(b) < 4+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint32(b[offset:4]), nil
}
func ReadB16(b []byte, offset int) (uint16, error) {
	if len(b) < 2+offset {
		return 0, io.EOF
	}
	return binary.BigEndian.Uint16(b[offset:2]), nil
}
func ReadB8(b []byte, offset int) (uint8, error) {
	if len(b) < 1+offset {
		return 0, io.EOF
	}
	return b[offset], nil
}

func ReadSB32(b []byte, offset int) (int32, error) {
	if len(b) < 4+offset {
		return 0, io.EOF
	}
	return int32(binary.BigEndian.Uint32(b[offset:4])), nil
}

// readSlice is like ReadBytes but returns a reference to internal buffer data.
func ReadSlice(b []byte, off int, lim int, delim byte) (line []byte, err error) {
	i := bytes.IndexByte(b[off:], delim)
	end := off + i + 1
	if i < 0 {
		end = len(b)
		err = io.EOF
	}
	line = b[off:end]
	return line, err
}

func WriteB64(b []byte, offset int, v uint64) error {
	if len(b) < 8+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint64(b[offset:8], v)
	return nil
}
func WriteB32(b []byte, offset int, v uint32) error {
	if len(b) < 4+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint32(b[offset:4], v)
	return nil
}
func WriteB16(b []byte, offset int, v uint16) error {
	if len(b) < 2+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint16(b[offset:2], v)
	return nil
}
func WriteB8(b []byte, offset int, v uint8) error {
	if len(b) < offset {
		return io.EOF
	}
	b[offset] = v
	return nil
}

func WriteSB32(b []byte, offset int, v int32) error {
	if len(b) < 4+offset {
		return io.EOF
	}
	binary.BigEndian.PutUint32(b[offset:4], uint32(v))
	return nil
}
