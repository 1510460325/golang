package main

import "fmt"

type Node struct {
	Val  string
	Next *Node
}

func (node Node) String() string {
	return fmt.Sprintf("node is %s", node.Val)
}

func (node *Node) hello() {
	fmt.Printf("hello, i'm a node %s", node.Val)
}

type Node1 struct {
	Val  string
	Next *Node1
}

func (node Node1) String() string {
	return fmt.Sprintf("node is %s", node.Val)
}

func (node *Node1) hello1() {
	fmt.Printf("hello, i'm node1 %s", node.Val)
}

func main() {
	var node fmt.Stringer
	node = &Node{}
	fmt.Printf("%T", node)
	//node = Node1{}
	if obj, ok := node.(*Node); ok {
		fmt.Println("node is a *Node")
		obj.hello()
	} else {
		fmt.Println("node is not a *Node")
	}
	switch node.(type) {
	case *Node:
		fmt.Println("this is node1")
		i := node.(*Node)
		i.hello()
		break
	case Node1:
		fmt.Println("this is node")
		i := node.(Node1)
		i.hello1()
		break
	}
}
