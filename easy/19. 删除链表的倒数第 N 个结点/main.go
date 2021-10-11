package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	// fast 先走n+1步
	for ; n >= 0; n-- {
		// 当fast提前走到终点 说明需移除第一个节点
		if fast == nil {
			return head.Next
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 此时slow.Next处于倒数第N个节点 摘除该节点 当倒数0时不需要移除
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return head
}
func main() {
	node6 := &ListNode{Val: 6, Next: nil}
	node5 := &ListNode{Val: 5, Next: node6}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	fmt.Println(removeNthFromEnd(node6, 1))
	for node1 != nil {
		fmt.Println(node1.Val)
		node1 = node1.Next
	}
}
