package experiment

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	b := []byte("ABC€")
	fmt.Println(b)

	s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(s)

	intArray := [5]int{1, 6, 9, 4, 3}
	for i, val := range intArray {
		fmt.Println("N:", i, " Value:", val)
	}

	for j, k := range "Golang" {
		fmt.Println(j, string(k))
	}

	// qq := [1..5]
	// for _, q := range [1..5] {
	// 	fmt.Println(q)
	// }
	// "github.com/bradfitz/iter"

	// Array
	var a1 = [5]int{1, 3, 5, 7, 9}
	var a2 = [...]int{2, 4, 6, 8, 10}
	var a3 = [4][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

	for j, k := range a1 {
		fmt.Printf("A1[%d]:%d\n", j, k)
	}

	for j, k := range a2 {
		fmt.Printf("A2[%d]:%d\n", j, k)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(a3[i][j], " ")
		}
		fmt.Println()
	}

	for _, i := range a3 {
		fmt.Println(i)
	}
}
