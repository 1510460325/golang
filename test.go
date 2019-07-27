package main

import (
	"./service"
	"fmt"
)

func main() {
	var a string = "abc"
	fmt.Println("hello, world", a)
	fmt.Println(UserService.Test(1, 2))
	fmt.Println(UserService.Id)
}
