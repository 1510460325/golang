package main

import "fmt"

func main() {
	fmt.Println(convert("sjyryjpyxiucuovurmrggjvnfzobqlkkwkbwtsvsa", 38))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	pic := [100000][100000]byte{{}}
	r, c := 0, 0
	direct := 1
	maxColumn := -1
	for i := 0; i < len(s); i++ {
		pic[r][c] = s[i]
		if maxColumn < c {
			maxColumn = c
		}
		if r == numRows-1 {
			direct = 2
		} else if r == 0 {
			direct = 1
		}
		switch direct {
		case 1:
			r++
			break
		case 2:
			r--
			c++
			break
		}
	}
	result := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j <= maxColumn; j++ {
			if pic[i][j] != 0 {
				result = append(result, pic[i][j])
			}
		}
	}
	return string(result)
}
