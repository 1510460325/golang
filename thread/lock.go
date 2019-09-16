package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	lock.Lock()
	go func() {
		defer wg.Done()
		lock.Lock()
		fmt.Println("协程2运行")
	}()
	fmt.Println("协程1运行")
	lock.Unlock()
	wg.Wait()
}
