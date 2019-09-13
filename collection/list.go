package main

import (
	"container/list"
	"fmt"
)

func find(names list.List, name string) *list.Element {
	for i := names.Front(); i != nil; i = i.Next() {
		if i.Value == name {
			return i
		}
	}
	return nil
}

func main() {
	var names *list.List = list.New()
	names.PushBack("1")
	names.PushFront("2")
	fmt.Println("遍历list")
	for i := names.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	target := find(*names, "1")
	fmt.Printf("查找并删除元素: %s\n", target.Value)
	names.Remove(target)
	fmt.Println("遍历list")
	for i := names.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
