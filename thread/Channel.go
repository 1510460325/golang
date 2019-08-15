package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var ch chan int

var wg sync.WaitGroup

func init() {
	ch = make(chan int)
	wg.Add(2)
}

func pop() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
		time.Sleep(500)
	}
}

func push() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(500)
		ch <- i
	}
}

func main() {
	go push()
	go pop()
	wg.Wait()
	inAndOut()
}

/*
	标准输入输出
*/
func inAndOut() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		str := in.Text()
		if str == "end" {
			break
		}
		fmt.Println(str)
	}
}
