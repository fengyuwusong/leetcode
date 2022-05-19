package main

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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

// reverse 反转区间[a, b)的元素， 注意左闭右开
func reverse(a, b *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur, next = a, a
	for cur != b {
		next = cur.Next
		// 翻转链表
		cur.Next = pre
		// 更新指针位置
		pre = cur
		cur = next
	}
	// 返回反转后的头结点
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间[a, b)包含k个待反转的元素
	a, b := head, head
	for i := 0; i < k; i++ {
		// 不足k个 不需要反转
		if b == nil {
			return head
		}
		b = b.Next
	}

	// 反转前k个元素
	newHead := reverse(a, b)
	// 递归反转后续链表并连接
	a.Next = reverseKGroup(b, k)
	return newHead
}

// @lc code=end
