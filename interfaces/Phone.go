package main

import "fmt"

type Phone interface {
	call()
}

type Nokia struct {
	id int
}

func (nokia *Nokia) func1() {
	fmt.Println("i'm fuc1")
}

func (nokia *Nokia) call() {
	nokia.func1()
	fmt.Println("i'm a Nokia")
}

func main() {
	fmt.Println("====接口实现====")
	var phone = Phone(new(Nokia))
	phone.call()
}
