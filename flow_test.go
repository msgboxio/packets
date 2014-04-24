package packets

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type S struct {
	a int
	b byte
	c int
}

var _ = fmt.Fprint
var _ = hex.Decode

func getBa(s *S) []byte {
	l := int(unsafe.Sizeof(*s))
	//a := unsafe.Alignof(*s)
	p := unsafe.Pointer(s)
	//fmt.Printf("len %d, alignment %d, pointer 0x%08x\n", l, a, p)

	_BYTE_SLICE := reflect.TypeOf([]byte(nil))
	header := &reflect.SliceHeader{uintptr(p), l, l}
	bs := reflect.NewAt(_BYTE_SLICE, unsafe.Pointer(header))
	return bs.Elem().Interface().([]byte)
}

func TestUnsafe(t *testing.T) {
	s1 := &S{0xffff, 0xaa, 0xdddd}
	s2 := &S{0xffff, 0xaa, 0xdddd}
	s3 := &S{0xffff, 0xab, 0xdddd}

	ba1 := getBa(s1)
	ba2 := getBa(s2)
	ba3 := getBa(s3)
	if !bytes.Equal(ba1, ba2) {
		t.Error("not equal")
	}
	if bytes.Equal(ba1, ba3) {
		t.Error("not unequal")
	}
	// fmt.Printf("%s", hex.Dump(ba1))
	// fmt.Printf("%s", hex.Dump(ba2))
	// fmt.Printf("%s", hex.Dump(ba3))
}
