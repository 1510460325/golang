package main

import "fmt"

func main() {
	fmt.Println(findMin2([]int{3, 3, 1, 3}))
}

func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		middle := (l + r) / 2
		if nums[middle] == nums[r] {
			min := nums[middle]
			for j := middle + 1; j <= r; j++ {
				if min > nums[j] {
					min = nums[j]
				}
			}
			if min < nums[middle] {
				return min
			} else {
				r = middle
			}
		} else if nums[middle] < nums[r] {
			r = middle
		} else {
			l = middle + 1
		}
	}
	return nums[l]
}
