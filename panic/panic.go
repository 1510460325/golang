package main

import (
	"errors"
	"fmt"
)

func except() {
	fmt.Println(recover())
}

func add(i, j int) int {
	return i + j
}

func main() {
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		fmt.Printf("recovered this error %v\n", e)
	}()
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if err, ok := e.(error); ok {
			fmt.Printf("recovered this error %v\n", err)
		} else {
			fmt.Println("i don't know how to do this panic, throw out!")
			panic("second panic")
		}
	}()
	panic(errors.New("12"))
}
