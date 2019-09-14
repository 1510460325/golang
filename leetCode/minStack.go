package main

import "container/list"

type MinStack struct {
	data *list.List
	min  *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{list.New(), list.New()}
}

func (this *MinStack) Push(x int) {
	if this.min.Front() == nil {
		this.min.PushFront(x)
	} else {
		this.min.PushFront(min(this.min.Front().Value.(int), x))
	}
	this.data.PushFront(x)
}

func (this *MinStack) Pop() {
	this.min.Remove(this.min.Front())
	this.data.Remove(this.data.Front())
}

func (this *MinStack) Top() int {
	return this.data.Front().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.min.Front().Value.(int)
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {

}
