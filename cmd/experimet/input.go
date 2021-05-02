package experiment

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestInput(t *testing.T) {
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
