package 剑指offer

import (
	"fmt"
	"testing"
)

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针解决
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	pre := head
	for pre != nil {
		temp := pre.Next
		pre.Next = cur
		cur = pre
		pre = temp
	}
	return cur
}

func TestReverseList(t *testing.T) {
	ln5 := &ListNode{Val: 5}
	ln4 := &ListNode{Val: 4, Next: ln5}
	ln3 := &ListNode{Val: 3, Next: ln4}
	ln2 := &ListNode{Val: 2, Next: ln3}
	ln1 := &ListNode{Val: 1, Next: ln2}
	d := reverseList(ln1)
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
