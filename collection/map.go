package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 1
	m[3] = 1
	val, ok := m[2]
	fmt.Printf("是否存在：%v, 值为:%d\n", ok, val)
	for key, val := range m {
		fmt.Printf("key is %d => %d\n", key, val)
	}
	delete(m, 1)
}
