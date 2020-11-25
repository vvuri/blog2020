package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log2file()
}

func linuxSyslog() {
	// LEVELS: debug, info, notice, warning, err, crit, alert, emerg
	// UNIX: syslogd, rsyslogd
	// /etc/rsyslog.conf

	// Only for Linux:
	// prgName := filepath.Base(os.Args[0])
	// sysLog, err := syslog.New(sysLog.LOG_INFO|syslog.LOG_LOCAL7, prgName)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.SetOutput(sysLog)
	// }
}

func logPanic() {
	log.Panic("Panic error\n")
	/*
		panic: Panic error

		goroutine 1 [running]:
		log.Panic(0xc000077f68, 0x1, 0x1)
				c:/go/src/log/log.go:351 +0xb7
		main.logPanic(...)
				C:/code/blog2020/goblog/cmd/experimet/logger.go:26
		main.main()
				C:/code/blog2020/goblog/cmd/experimet/logger.go:22 +0x65
		exit status 2
	*/
}

func logFatal() {
	log.Fatal("Fatal error\n")
	/*
		2020/11/25 11:21:43 Fatal error
		exit status 1
	*/
}

func log2file() {
	var LOGFILE = "logger.log"

	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	ilog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	ilog.SetFlags(log.LstdFlags)

	ilog.Println("Hello")
	ilog.Println("Second lint")

	/*
		customLogLineNumber 2020/11/25 11:44:23 Hello
		customLogLineNumber 2020/11/25 11:44:23 Second lint
	*/
}
