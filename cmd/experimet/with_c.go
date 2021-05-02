package experiment

// #include <stdio.h>
// void call_c() {
//   printf("Calling C code. \n")
// }
//-import "C"
import (
	"fmt"
	"testing"
)

func TestWithC(t *testing.T) {
	fmt.Println("Run C code in Go programm:")
//-	C.call_c()
}

// Need GCC
