package main

import (
	"fmt"
	"unsafe"
)

type Bad struct {
	A byte  // 1 байт
	B int64 // 8 байт
	C byte  // 1 байт
	D int32 // 4 байта
}
type Good struct {
	B int64 // 8 байт
	D int32 // 4 байта
	A byte  // 1 байт
	C byte  // 1 байт
}

func main() {

	fmt.Println("Size of Bad:", unsafe.Sizeof(Bad{}))
	fmt.Println("Size of Good:", unsafe.Sizeof(Good{}))
}
