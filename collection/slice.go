package main

import "fmt"

func diff(arr1 []int, arr2 []int) bool {
	if len(arr1) == len(arr2) {
		return true
	}
	fmt.Println(len(arr1))
	return false
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(arr[1:3][:6])
	fmt.Println("len is : ", len(arr[2:5])) // 3
	fmt.Println("cap is : ", cap(arr[2:5])) // 5
	fmt.Println("len is : ", len(arr[3:5])) // 2
	fmt.Println("cap is : ", cap(arr[3:5])) // 4

	arr2 := [...]int{0, 0, 0, 0, 0, 0}
	copy(arr[:], arr2[:])
	fmt.Println(arr)
	fmt.Println(arr2)
}
