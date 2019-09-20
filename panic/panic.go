package main

import "fmt"

func except() {
	fmt.Println(recover())
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	panic(12)
}
