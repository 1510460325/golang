package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func println2(str string, n int) {
	defer func() {
		fmt.Println("=====thread end =====")
		waitGroup.Done()
	}()
	for i := 1; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	waitGroup.Add(2)
	go println2("aaa", 5)
	go println2("bbb", 5)
	waitGroup.Wait()
}
