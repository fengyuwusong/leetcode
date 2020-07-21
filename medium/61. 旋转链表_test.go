package medium

import (
	"testing"
)

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 使用数组记录链表节点
// 解法2 使用快慢指针即可 时间复杂度O(n) 空间复杂度O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	// 增加代码健壮性
	if head == nil {
		return nil
	}
	// k==0直接返回
	if k == 0 {
		return head
	}
	fast, slow := head, head
	length := 0
	// fast指针先走k节点
	for i := 0; i < k; i++ {
		length++
		if fast.Next == nil {
			fast = head
			i = 0
			k = k%length + 1
		} else {
			fast = fast.Next
		}
	}
	// fast==head 说明不用旋转
	if fast == head {
		return head
	}
	// fast slow同时遍历 fast到达终点时停止遍历
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 此时slow节点为新链表终点 slow.next节点为新链表head fast.next连接原链表head
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

func TestRotateRight(t *testing.T) {
	//ln5 := &ListNode{Val: 5}
	//ln4 := &ListNode{Val: 4, Next: ln5}
	ln3 := &ListNode{Val: 2, Next: nil}
	ln2 := &ListNode{Val: 1, Next: ln3}
	ln1 := &ListNode{Val: 0, Next: ln2}
	rotateRight(ln1, 4)
	//rotateRight(ln5, 1)
}
