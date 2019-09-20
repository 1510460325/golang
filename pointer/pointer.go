package main

import "fmt"

func change(ptr *int) {
	fmt.Printf("指针地址：%v, 值：%v\n", &ptr, *ptr)
}

func main() {
	a := 12
	// 新建一个指针
	var ptr *int = new(int)
	fmt.Printf("指针地址：%v, 值：%v\n", &ptr, *ptr)
	// 指针赋值
	ptr = &a
	fmt.Printf("指针地址：%v, 值：%v\n", &ptr, *ptr)
	change(ptr)
	fmt.Printf("指针地址：%v, 值：%v\n", &ptr, *ptr)
}
