package main

import (
	"fmt"
)

func main() {
	var s string = "  "
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	time := make(map[byte]int)
	start, end := 0, 1
	time[s[start]] = 1
	maxLen := -1
	for ; end < len(s); end++ {
		_, ok := time[s[end]]
		if !ok {
			time[s[end]] = 0
		}
		time[s[end]]++
		for time[s[end]] == 2 {
			time[s[start]]--
			start++
		}
		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
	}
	return maxLen
}
