package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No float args.")
		os.Exit(1)
	}

	arguments := os.Args

	for i := 1; i < len(arguments); i++ {
		min, _ := strconv.ParseFloat(arguments[i], 64)

		fmt.Printf("arg[%d]: %f\n", i, min)
	}
}
