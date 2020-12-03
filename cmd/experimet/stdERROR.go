package main

import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Message to stdout\n")
	io.WriteString(os.Stderr, "Message to stderr\n")
	// go run ./cmd/experimet/stdERROR.go 2>error
}
