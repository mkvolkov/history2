package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Git Replace demo")

	var a, b = 5, 7

	fmt.Printf("5 + 7 = %d\n", a+b)
	fmt.Printf("5 - 7 = %d\n", a-b)
	fmt.Printf("5 * 7 = %d\n", a*b)
	fmt.Printf("5 / 7 = %d\n", a/b)

	fmt.Printf("Size of int = %d\n", unsafe.Sizeof(a))

	var x float64 = 7.62
	fmt.Printf("Size of float64 = %d\n", unsafe.Sizeof(x))
	fmt.Printf("Size of float32 = %d\n", unsafe.Sizeof(float32(x)))
}
