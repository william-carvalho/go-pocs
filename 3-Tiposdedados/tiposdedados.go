package main

import "fmt"

func main() {
	int8, int16, int32, int64 := int8(10), int16(20), int32(30), int64(40)
	uint8, uint16, uint32, uint64 := uint8(10), uint16(20), uint32(30), uint64(40)
	float32, float64 := float32(10.5), float64(20.99)
	var complex64 complex64 = complex(1, 2)
	var complex128 complex128 = complex(2, 3)
	str := "String de exemplo"
	b := true
	var byteVar byte = 255
	var runeVar rune = 'a'

	fmt.Println(int8, int16, int32, int64)
	fmt.Println(uint8, uint16, uint32, uint64)
	fmt.Println(float32, float64)
	fmt.Println(complex64, complex128)
	fmt.Println(str)
	fmt.Println(b)
	fmt.Println(byteVar)
	fmt.Println(runeVar)

}
