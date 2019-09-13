package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{1, 2, 3}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		middle := (l + r) / 2
		if nums[middle] <= nums[r] {
			r = middle
		} else {
			l = middle + 1
		}

	}
	return nums[l]
}
