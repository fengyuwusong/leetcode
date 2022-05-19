package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

func printListNode(head *ListNode) {
	if head != nil {
		fmt.Printf("%v->", head.Val)
		printListNode(head.Next)
		return
	}
	fmt.Printf("\n")
}

func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{3, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	head := ListNode{1, &node2}
	deleteDuplicates(&head)
	printListNode(&head)
}
