package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	val  *TreeNode
	next *BSTIterator
}

func Constructor1(root *TreeNode) BSTIterator {
	ans := &BSTIterator{}
	tail := ans
	stack := list.New()
	now := root
	if now == nil {
		return *ans
	}
	for now != nil || stack.Len() > 0 {
		for now != nil {
			stack.PushFront(now)
			now = now.Left
		}
		now = stack.Remove(stack.Front()).(*TreeNode)
		tail.next = new(BSTIterator)
		tail = tail.next
		tail.val, tail.next = now, nil
		now = now.Right
	}
	return *ans.next
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	val := this.val
	if val == nil {
		return -1
	}
	if this.next == nil {
		this.val = nil
	} else {
		this.val = this.next.val
		this.next = this.next.next
	}
	return val.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.val != nil
}

func main() {
	root := &TreeNode{7, nil, nil}
	root.Left = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Right.Left = &TreeNode{9, nil, nil}
	root.Right.Right = &TreeNode{20, nil, nil}
	iterator := Constructor1(root)
	println(iterator.Next())    // 返回 3
	println(iterator.Next())    // 返回 7
	println(iterator.HasNext()) // 返回 true
	println(iterator.Next())    // 返回 9
	println(iterator.HasNext()) // 返回 true
	println(iterator.Next())    // 返回 15
	println(iterator.HasNext()) // 返回 true
	println(iterator.Next())    // 返回 20
	println(iterator.HasNext()) // 返回 false
}
