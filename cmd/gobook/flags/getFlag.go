package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say ...")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hi %s!\n", *name)
	}
}
