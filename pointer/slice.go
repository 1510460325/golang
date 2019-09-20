package main

import "fmt"

func print(a []int, b []int) {
	fmt.Println("=====")
	fmt.Printf("cap：%v, 数组取值:%v\n", cap(a), a)
	fmt.Printf("cap：%v, 数组取值:%v\n", cap(b), b)
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums1 := nums[0:2]
	print(nums, nums1)
	//nums1[0] = 12
	nums1 = append(nums1, 12, 12)
	print(nums, nums1)
	nums1 = append(nums1, 12, 12)
	print(nums, nums1)
}
