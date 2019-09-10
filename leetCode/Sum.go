package main

import "fmt"

func main() {
	var ans []int = twoSum([]int{2, 7, 11, 15}, 13)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {
	ans := []int{-1, -1}
	visit := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		val, ok := visit[target-nums[i]]
		if ok {
			ans[0], ans[1] = val, i
			break
		}
		visit[nums[i]] = i
	}
	return ans
}
