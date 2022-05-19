package main

import "fmt"

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
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

var successor *ListNode // 后驱节点

// 反转前n个链表节点
func reverseN(head *ListNode, n int) *ListNode {
	// 递归出口 同时记录后驱节点 用于反转后重新引用
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 重置后驱节点
	head.Next = successor
	return last
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		// 相当于反转前 right 个元素
		return reverseN(head, right)
	}

	// 步进left后反转前right个元素 步进right也需要减少
	head.Next = reverseBetween(head.Next, left-1, right-1)

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

// @lc code=end
func main() {
	node5 := ListNode{Val: 5}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	head := ListNode{
		Val:  1,
		Next: &node2,
	}

	printListNode(&head)
	head2 := reverseBetween(&head, 2, 4)
	printListNode(head2)
}
