package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func println2(str string, n int) {
	defer fmt.Println("=====thread end =====")
	defer waitGroup.Done()
	for i := 1; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	waitGroup.Add(2)
	go println2("aaa", 10)
	go println2("bbb", 10)
	waitGroup.Wait()
}
