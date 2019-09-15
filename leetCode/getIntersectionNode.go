package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for {
		if q == p {
			return q
		}
		if q == nil || p == nil {
			return nil
		}
		if p.Next == nil {
			p = headB
			headB = nil
		} else {
			p = p.Next
		}
		if q.Next == nil {
			q = headA
			headA = nil
		} else {
			q = q.Next
		}
	}
}
