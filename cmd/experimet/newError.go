package experiment

import (
	"errors"
	"log"
	"testing"
)

func returnError(a, b int) error {
	if a == b {
		return errors.New("Error: A=B")
	} else {
		return nil
	}
}

func TestError(t *testing.T) {
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
