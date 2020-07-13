package easy

import (
	"fmt"
	"testing"
)

/**
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		if cur.Val == val {
			cur = cur.Next
			if pre == nil {
				head = cur
			} else {
				pre.Next = cur
			}
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head
}

func TestRemoveElements(t *testing.T) {
	ln5 := &ListNode{Val: 5}
	ln4 := &ListNode{Val: 4, Next: ln5}
	ln3 := &ListNode{Val: 3, Next: ln4}
	ln2 := &ListNode{Val: 2, Next: ln3}
	ln1 := &ListNode{Val: 1, Next: ln2}
	d := removeElements(ln1, 1)
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
