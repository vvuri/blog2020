package experiment

import (
	"fmt"
	"runtime"
	"testing"
)

func getSystemParam() {
	fmt.Println("Runtime:", runtime.Compiler)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num Goroutines:", runtime.NumGoroutine())
}

func TestEnv(t *testing.T) {
	getSystemParam()
}
