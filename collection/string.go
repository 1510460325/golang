package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "abc你好!"
	for _, ch := range []rune(str) {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c", ch)
		bytes = bytes[size:]
	}
}
