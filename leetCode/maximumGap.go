package main

import "sort"

func main() {

}

func maximumGap(nums []int) int {
	max := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		gap := nums[i] - nums[i-1]
		if gap < 0 {
			gap *= -1
		}
		if max < gap {
			max = gap
		}
	}
	return max
}
