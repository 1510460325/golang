package main

import (
	"./service"
	"fmt"
)

type Person struct {
	id   int
	name string
}

func main() {
	// := 赋值并申明
	b := UserService.Test(1, 2)
	//fmt.Println(UserService.Test(1, 2))
	//fmt.Println(UserService.Id)
	// if else 语句
	fmt.Println("====if 判断====")
	if b > 2 {
		fmt.Println(b)
	} else {
		fmt.Println(2)
	}
	// for循环
	fmt.Println("====for 循环====")
	n := 5
	i := 1
	for i < n {
		fmt.Println("当前运行到", i)
		i++
	}
	// 函数闭包
	fmt.Println("==== 函数闭包 ====")
	fmt.Println(getI()())
	// 数组
	fmt.Println("======数组======")
	var bags [15]int
	// 数组赋值
	bags = [15]int{15, 15}
	fmt.Println(bags[1])
	fmt.Println("====数组切片====")
	fmt.Println(bags[1:4])
	// 指针
	fmt.Println("====指针====")
	var ptr = &b
	fmt.Printf("指针地址：%x\n", ptr)
	fmt.Println("====结构体====")
	var person = Person{id: 1, name: "122"}
	fmt.Println(person)
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
