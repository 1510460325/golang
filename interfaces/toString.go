package main

import "fmt"

type Node struct {
	Val  string
	Next *Node
}

func (node *Node) String() string {
	return fmt.Sprintf("node is %s", node.Val)
}

func main() {
	node := new(Node)
	node1 := Node{}
	fmt.Println(node)
	fmt.Println(node1)
}
