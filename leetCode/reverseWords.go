package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " the sky is blue "
	fmt.Println(reverseWords(str))
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	index := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			index = i
			break
		}
	}
	if index != -1 {
		pre := string([]byte(s)[:index])
		now := string([]byte(s)[index+1:])
		return now + " " + reverseWords(pre)
	}
	return s
}
