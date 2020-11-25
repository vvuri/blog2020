package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Println(len(arguments))

	if len(arguments) > 1 {
		fmt.Println(arguments[1])
	}

	n, err := io.WriteString(os.Stdout, "test_stdout")

	if err != nil {
		panic(err)
	}

	fmt.Printf("n: %d\n", n)
}
