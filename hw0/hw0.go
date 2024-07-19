package main

import (
	"fmt"
)

func main() {
	var (
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		f32  float32
		f64  float64
		b    bool
		r    rune
		by   byte
		s    string
		c64  complex64
		c128 complex128
	)

	fmt.Println("Default value")
	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)
	fmt.Println("uint:", ui)
	fmt.Println("uint8:", ui8)
	fmt.Println("uint16:", ui16)
	fmt.Println("uint32", ui32)
	fmt.Println("uint64", ui64)
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)
	fmt.Println("bool:", b)
	fmt.Println("rune:", r)
	fmt.Println("byte:", by)
	fmt.Println("string", s)
	fmt.Println("complex64", c64)
	fmt.Println("complex128", c128)

	i = -1
	i8 = -8
	i16 = -16
	i32 = -32
	i64 = -64
	ui = 1
	ui8 = 8
	ui16 = 16
	ui32 = 32
	ui64 = 64
	f32 = 3.14
	f64 = 2.715642
	b = true
	r = 'A'
	by = 255
	s = "Hello World"
	c64 = 1 + 2i
	c128 = 2 + 3i
	fmt.Println("Changed Values:")
	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)
	fmt.Println("uint:", ui)
	fmt.Println("uint8:", ui8)
	fmt.Println("uint16:", ui16)
	fmt.Println("uint32", ui32)
	fmt.Println("uint64", ui64)
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)
	fmt.Println("bool:", b)
	fmt.Println("rune:", r)
	fmt.Println("byte:", by)
	fmt.Println("string", s)
	fmt.Println("complex64", c64)
	fmt.Println("complex128", c128)
}
