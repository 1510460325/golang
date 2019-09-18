package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(-50, 8))
}

func fractionToDecimal(number, demomin int) string {
	pre := ""
	if number*demomin < 0 {
		pre += "-"
	}
	if number < 0 {
		number *= -1
	}
	if demomin < 0 {
		demomin *= -1
	}
	a, b := getDivideAndComplete(number, demomin)
	pre += strconv.Itoa(a)
	if b == 0 {
		return pre
	}
	pre += "."

	next := ""
	visit := make(map[int]int)
	flag, count, start := false, -1, -1
	for {
		count++
		b *= 10
		index, ok := visit[b]
		if !ok {
			visit[b] = count
		} else {
			flag = true
			start = index
			break
		}
		a, b = getDivideAndComplete(b, demomin)
		next += strconv.Itoa(a)
		if b == 0 {
			break
		}
	}
	if flag {
		next = next[:start] + "(" + next[start:] + ")"
	}
	return pre + next
}

func getDivideAndComplete(a, b int) (int, int) {
	return a / b, a % b
}
