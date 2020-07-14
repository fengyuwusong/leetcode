package 程序员面试金典

import (
	"fmt"
	"testing"
)

/**
编写一个函数，检查输入的链表是否是回文的。



示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true


进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针指针解法 时间复杂度O(n) 空间复杂度O(1)
// 1. 快慢指针确定链表中点
// 2. 翻转后半部链表  如fast==nil&&fast.next=nil 说明链表长度单数 则从slow.next 开始反转 反之从slow开始
// 3. 依次对比前后链表
// 4. 如要求不破坏原链表则翻转后半部链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 确定链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 链表长度奇数 从中点的下一个节点开始
	if fast != nil {
		slow = slow.Next
	}
	// 反转链表
	var slowCur *ListNode
	for slow != nil {
		temp := slow.Next
		slow.Next = slowCur
		slowCur = slow
		slow = temp
	}
	// 对比反转后的前后链表
	for slowCur != nil {
		if slowCur.Val != head.Val {
			return false
		}
		slowCur = slowCur.Next
		head = head.Next
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	ln5 := &ListNode{Val: 5}
	ln4 := &ListNode{Val: 4, Next: ln5}
	ln3 := &ListNode{Val: 3, Next: ln4}
	ln2 := &ListNode{Val: 4, Next: ln3}
	ln1 := &ListNode{Val: 5, Next: ln2}
	fmt.Println(isPalindrome(ln1))
}
