package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func strangePrinter(str string) int {
	if str == "" {
		return 0
	}
	dp := [101][101]int{{}}
	for i := 0; i < len(str); i++ {
		dp[i][i] = 1
	}
	for i := len(str) - 2; i >= 0; i-- {
		for j := i + 1; j < len(str); j++ {
			dp[i][j] = dp[i+1][j] + 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][k]+dp[k+1][j], dp[i][j])
			}
			if str[i] == str[j] {
				dp[i][j] = min(dp[i+1][j], dp[i][j])
			}
		}
	}
	return dp[0][len(str)-1]
}

func main() {
	fmt.Println(strangePrinter("tbgtgb"))
}
