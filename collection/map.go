package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 1
	m[3] = 1
	for key, val := range m {
		fmt.Printf("key is %d => %d\n", key, val)
	}
}
