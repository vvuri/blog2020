package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	fmt.Println("Size int:", unsafe.Sizeof(i))

	array := [...]int{3, 1, -5, -4, 6}
	pointer := &array[0]
	fmt.Println(*pointer, pointer)

	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println(memoryAddress)
	fmt.Println("Size one element:", unsafe.Sizeof(array[0]))

	for i := 0; i < len(array); i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Println(*pointer, pointer)
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
}
