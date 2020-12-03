package main

import (
	"fmt"
	"runtime"
)

func getSystemParam() {
	fmt.Println("Runtime:", runtime.Compiler)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num Goroutines:", runtime.NumGoroutine())
}

func main() {
	getSystemParam()
}
