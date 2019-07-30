package main

import (
	"fmt"
	"time"
)

func println2(str string, n int) {
	for i := 1; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	go println2("bweqe", 10)
	println2("aqwee", 10)
}
