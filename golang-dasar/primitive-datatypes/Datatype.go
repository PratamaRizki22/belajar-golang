package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var a int = 10
	var b uint8 = 11
	var c uint16 = 12
	var d uint32 = 13
	var e uint64 = 14
	var f1 int8 = 15
	var g1 int16 = 16
	var f2 int32 = 17
	var g2 int64 = 18
	var h float32 = 19.5
	var i float64 = 10.5
	var complex1 complex64 = 2 + 3i
	complex2 := complex(10, 13)
	complex3 := complex128(10 + 1i)
	var huruf string = "koko"
	var huruf2 string = "kokokokokoko"

	fmt.Printf("a: value=%v, type=%T, size=%d bytes, max=%d\n", a, a, unsafe.Sizeof(a), int64(math.MaxInt))
	fmt.Printf("b: value=%v, type=%T, size=%d bytes, max=%d\n", b, b, unsafe.Sizeof(b), uint8(math.MaxUint8))
	fmt.Printf("c: value=%v, type=%T, size=%d bytes, max=%d\n", c, c, unsafe.Sizeof(c), uint16(math.MaxUint16))
	fmt.Printf("d: value=%v, type=%T, size=%d bytes, max=%d\n", d, d, unsafe.Sizeof(d), uint32(math.MaxUint32))
	fmt.Printf("e: value=%v, type=%T, size=%d bytes, max=%d\n", e, e, unsafe.Sizeof(e), uint64(math.MaxUint64))
	fmt.Printf("f1: value=%v, type=%T, size=%d bytes, max=%d\n", f1, f1, unsafe.Sizeof(f1), int8(math.MaxInt8))
	fmt.Printf("g1: value=%v, type=%T, size=%d bytes, max=%d\n", g1, g1, unsafe.Sizeof(g1), int16(math.MaxInt16))
	fmt.Printf("f2: value=%v, type=%T, size=%d bytes, max=%d\n", f2, f2, unsafe.Sizeof(f2), int32(math.MaxInt32))
	fmt.Printf("g2: value=%v, type=%T, size=%d bytes, max=%d\n", g2, g2, unsafe.Sizeof(g2), int64(math.MaxInt64))
	fmt.Printf("h: value=%v, type=%T, size=%d bytes, max=%g\n", h, h, unsafe.Sizeof(h), float32(math.MaxFloat32))
	fmt.Printf("i: value=%v, type=%T, size=%d bytes, max=%g\n", i, i, unsafe.Sizeof(i), float64(math.MaxFloat64))
	fmt.Printf("complex1: value=%v, type=%T, size=%d bytes\n", complex1, complex1, unsafe.Sizeof(complex1))
	fmt.Printf("complex2: value=%v, type=%T, size=%d bytes\n", complex2, complex2, unsafe.Sizeof(complex2))
	fmt.Printf("complex3: value=%v, type=%T, size=%d bytes\n", complex3, complex3, unsafe.Sizeof(complex3))
	fmt.Printf("huruf1: value=%v, type=%T, size=%d bytes\n", huruf, huruf, unsafe.Sizeof(huruf))
	fmt.Printf("huruf2: value=%v, type=%T, size=%d bytes\n", huruf2, huruf2, unsafe.Sizeof(huruf2))

	fmt.Printf("\nAdditional formatting:\n")
	fmt.Printf("a in binary: %b\n", a)
	fmt.Printf("b as character: %c\n", b)
	fmt.Printf("d in hexadecimal: %x\n", d)
	fmt.Printf("h in scientific notation: %e\n", h)
	fmt.Printf("i as quoted string: %q\n", fmt.Sprintf("%f", i))
}
