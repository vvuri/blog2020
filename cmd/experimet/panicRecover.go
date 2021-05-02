package experiment

import (
	"fmt"
	"testing"
)

func a() {
	fmt.Println("Run A")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Panic:", c)
			fmt.Println("Recover A")
		}
	}()
	fmt.Println("Call B")
	b()
	fmt.Println("After call B")
}

func b() {
	fmt.Println("Run B")
	panic("Panic in B")
	fmt.Println("End B")
}

func TestPanic(t *testing.T) {
	a()
	fmt.Println("End main")
}
