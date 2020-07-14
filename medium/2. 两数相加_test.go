package medium

import (
	"fmt"
	"testing"
)

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addFlag := 0
	ans := &ListNode{Next: nil}
	cur := ans
	for l1 != nil || l2 != nil || addFlag == 1 {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + addFlag
		node := &ListNode{Val: sum % 10}
		cur.Next = node
		cur = node

		addFlag = 0
		if sum >= 10 {
			addFlag = 1
		}
	}
	return ans.Next
}

func TestAddTwoNumbers(t *testing.T) {
	ln4 := &ListNode{Val: 1}
	ln2 := &ListNode{Val: 9}
	ln1 := &ListNode{Val: 9, Next: ln2}
	d := addTwoNumbers(ln1, ln4)
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
