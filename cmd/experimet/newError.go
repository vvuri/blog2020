package main

import (
	"errors"
	"log"
)

func returnError(a, b int) error {
	if a == b {
		return errors.New("Error: A=B")
	} else {
		return nil
	}
}

func main() {
	// err := returnError(1, 2)
	// fmt.Println(err)

	err := returnError(1, 1)
	if err != nil {
		// defer os.Exit(10)
		// log.Println(err)
		log.Panic(err)
		// panic(err)
	}
}
