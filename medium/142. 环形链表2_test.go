package medium

import (
	"fmt"
	"testing"
)

/**
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。




进阶：
你是否可以不用额外空间解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法1 哈希表 时间复杂度O(n) 空间复杂度O(n)
// 解法2 双指针 时间复杂度O(n) 空间复杂度O(1)
// 双指针slow fast 分别以步进 1 2的速度前行
// 那么等两者相遇时 假设slow遍历的节点个数 = x + y + i*(n)  fast遍历的节点个数 = x + y + j*(n)
// (x=达到环起始节点个数 y=在环中相遇的节点距离环开始的个数 z=相遇节点到环起始点的节点个数 n代表环节点的个数=y+z
// i代表慢指针遍历的圈数 j代表快指针遍历环的圈数 j>i)
// 由两者速率当相遇时可得公式 2*(x + y + i*(n)) = x+y+j*(n) => x+y = (j-2i)*n = (j-2i)*(y+z)
// => x = (j-2i)*(y+z)-y
// 由题意可知 j>=2i必定成立
// 当j-2i=1时 x = z
// 当j-2i=o>1时 x = (o-1)(y+z)+y+z-y = (o-1)(y+z)+z
// 即遍历1/o-1圈后走过x必定到达z相遇点
// 那么我们仅需设立第三个指针从头开始，与快慢指针相遇点一起遍历 那么当遍历1/o-1个环后再经过x/z距离必定到达环起点
// 如图
// 1->2->3->4
//    |     |
//    -------
// 快慢指针将会在4节点相遇 x=1 y=2 z=1 n=y+z=3
// 当重新设置第三个指针从1开始以速率1遍历 那么将会在相遇节点4完成1圈环后在2节点相遇 => x=z
// 解法2
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 双指针遍历到两指针相遇
	slot, fast := head, head
	for fast != nil && fast.Next != nil {
		slot, fast = slot.Next, fast.Next.Next
		if slot == fast {
			slot = head
			for slot != fast {
				slot = slot.Next
				fast = fast.Next
			}
			return slot
		}
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	ln1 := &ListNode{Val: 3}
	ln2 := &ListNode{Val: 2}
	ln3 := &ListNode{Val: 0}
	ln4 := &ListNode{Val: -4}
	ln1.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln2
	fmt.Println(detectCycle(ln1).Val)
}
