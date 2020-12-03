package main

// #include <stdio.h>
// void call_c() {
//   printf("Calling C code. \n")
// }
import "C"
import "fmt"

func main() {
	fmt.Println("Run C code in Go programm:")
	C.call_c()
}

// Need GCC
