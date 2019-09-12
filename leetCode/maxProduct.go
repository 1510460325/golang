package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	imax, imin := 1, 1
	maxAns := -65535 * 65535
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tem := imax
			imax = imin
			imin = tem
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		maxAns = max(maxAns, imax)
	}
	return imax
}
