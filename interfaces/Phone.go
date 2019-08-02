package main

import "fmt"

type Phone interface {
	call()
}

type Nokia struct {
}

func (nokia Nokia) call() {
	fmt.Println("i'm a Nokia")
}

func (nokia Nokia) call1() {
	fmt.Println("i'm a Nokia1")
}

func main() {
	fmt.Println("====接口实现====")
	var phone = Phone(new(Nokia))
	phone.call()
}
