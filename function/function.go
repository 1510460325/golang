package main

import "fmt"

type Pow func(int) int

func invoker(f Pow) {
	fmt.Println(f(2))
}

type MyInt = int

func main() {
	var function = func(i int) int {
		return i * i
	}
	fmt.Printf("the type is : %T\n", function)
	invoker(function)

	var i MyInt = 1
	fmt.Printf("the type is : %T\n", i)
}
