package main

import "fmt"

type Base struct {
}

type Function interface {
	func1()
}

func (base Base) func1() {
	fmt.Println("i'm a Base func1")
}

type Child struct {
	Base
}

func (child Child) func1() {
	fmt.Println("i'm a Child func1")
}

func main() {
	var function Function = new(Base)
	function.func1()
	//var base Base = new(Child)
	// 父结构体指向子结构体会报错
}
