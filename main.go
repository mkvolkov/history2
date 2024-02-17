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

	sl1 := make([]int, 6)
	sl1[2] = 700

	fmt.Println(sl1)
	fmt.Println("size of sl1 = ", unsafe.Sizeof(sl1))
	fmt.Println("length of sl1 = ", len(sl1))
	fmt.Println("capacity of sl1 = ", cap(sl1))

	mp1 := make(map[string]int)
	mp1["Arnold Schwarzenegger"] = len("Arnold Schwarzenegger")
	mp1["Max Payne"] = len("Max Payne")
	mp1["Daniil Dankovsky"] = len("Daniil Dankovsky")

	fmt.Println("size of mp1 = ", unsafe.Sizeof(mp1))
	fmt.Println("length of mp1 = ", len(mp1))
}
