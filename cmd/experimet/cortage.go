package experiment

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"
	"time"
)

func retThree(x int) (int, int, int) {
	return x * x, x * 2, -x
}

func regOne(line string) {
	r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
	if r1.MatchString(line) {
		match := r1.FindStringSubmatch(line)
		fmt.Print(match)
		d1, _ := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
		fmt.Println(d1)
	}
}

func stdOutLog(myString string) {
	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}

func TestCortage(t *testing.T) {
	fmt.Println(retThree(5))
	x1, x2, _ := retThree(3)
	fmt.Println(x1, x2)

	fmt.Print("Regexp: ")
	regOne("21/Nov/2017:19:28:09 +0200")

	stdOutLog("Log to stdout")
}
