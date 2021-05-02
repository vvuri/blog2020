package experiment

import (
	"fmt"
	s "strings"
	"testing"
	"unicode"
)

// Рунный литерал — это символ, заключенный в одинарные кавычки.
func runRune() {
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)
}

// Строка Go — это байтовый срез, предназначенный только для чтения
func str() {
	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)
	fmt.Printf("s3 length: %d\n", len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}

func runUnicode() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}

func str2() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
}

func str3() {

	f("%s\n", s.Split("abcd efg", ""))
}

var f = fmt.Printf

func TestString(t *testing.T) {
	// runRune()
	// str()
	// runUnicode()
	// str2()
	str3()
}
