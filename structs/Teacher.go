package main

import (
	"./base"
	"fmt"
)

type Teacher struct {
	base.User
	TId string
}

func (teacher Teacher) introduce() {
	teacher.User.SayHello()
	fmt.Println("i'm a teacher, ID is " + teacher.TId)
}

func main() {
	a := Teacher{base.User{"***", "wangzy", 111}, "111"}
	a.introduce()
}
