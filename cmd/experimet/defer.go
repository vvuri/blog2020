package main

import "fmt"

func main() {
	for i := 3; i > 0; i-- {
		defer fmt.Printf("%d ", i)
	}

	defer fmt.Println()

	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}

	defer fmt.Println()

	for i := 3; i > 0; i-- {
		defer func(k int) {
			fmt.Printf("%d ", k)
		}(i)
	}

}
