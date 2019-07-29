package main

import (
	"./service"
	"fmt"
)

func main() {
	// := 赋值并申明
	a := "abc"
	b := UserService.Test(1, 2)
	fmt.Println("hello, world", a)
	fmt.Println(UserService.Test(1, 2))
	fmt.Println(UserService.Id)
	// if else 语句
	if b > 2 {
		fmt.Println(b)
	} else {
		fmt.Println(2)
	}
	// for循环
	n := 5
	i := 1
	for i < n {
		fmt.Println("当前运行到", i)
		i++
	}
	// 函数闭包
	fmt.Println(getI()())
}

/*
  函数闭包：返回的是一个函数，再次调用则返回结果
*/
func getI() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
