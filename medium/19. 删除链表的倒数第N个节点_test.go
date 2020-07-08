package medium

import (
	"fmt"
	"testing"
)

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针解法 先使用快指针走n步 然后快慢指针相同速率走 删除慢指针节点即可
// 特殊情况可用哨兵节点解决 如删除头节点

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if n >= 0 {
			n--
			fast = fast.Next
		} else {
			slow, fast = slow.Next, fast.Next
		}
	}
	// 大于0说明仅一个节点 返回nil
	if n > 0 {
		return nil
	} else if n == 0 {
		// 等于0说明删除头节点 返回头节点的下一个节点
		return head.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func TestRemoveNthFormEnd(t *testing.T) {
	ln5 := &ListNode{Val: 5}
	ln4 := &ListNode{
		Val:  4,
		Next: ln5,
	}
	ln3 := &ListNode{
		Val:  3,
		Next: ln4,
	}
	ln2 := &ListNode{
		Val:  2,
		Next: ln3,
	}
	ln1 := &ListNode{
		Val:  1,
		Next: ln2,
	}
	removeNthFromEnd(ln1, 2)
	for ln1 != nil {
		fmt.Println(ln1.Val)
		ln1 = ln1.Next
	}
	test2 := &ListNode{Val: 2}
	test1 := &ListNode{Val: 1, Next: test2}

	n := removeNthFromEnd(test1, 2)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
