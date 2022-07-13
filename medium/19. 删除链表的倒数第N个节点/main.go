package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{0, head}
	var fast, slow = dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 此时slow指向倒数第n-1个节点
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	var head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = removeNthFromEnd(head, 2)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
