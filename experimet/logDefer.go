package main

import (
	"log"
	"os"
)

var LOGFILE = "logDefer.log"

func one(aLog *log.Logger) {
	aLog.Println("---Begin---")
	defer aLog.Println("---End-----")

	for i := 0; i < 5; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("---Begin---")
	defer aLog.Println("---End-----")

	for i := 5; i > 0; i-- {
		aLog.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println(">> Start")
	defer iLog.Println(">> Stop")
	one(iLog)
	two(iLog)
}
