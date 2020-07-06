package easy

import (
	"fmt"
	"testing"
)

/**
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。




进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1 哈希表 时间复杂度O(n) 空间复杂度O(1)
// 解法2 双指针 通过使用不同素的的快、慢两个指针遍历链表。
// 如果列表中不存在环，最终快指针会最先到达尾部，如存在环，快指针会追上慢指针。
// 时间复杂度O(N+K) = O(n) 空间复杂度O(1)

// 解法2
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		// 快指针为空时不存在闭环
		if fast == nil || fast.Next == nil {
			return false
		}
		// 慢指针向前移动1
		slow = slow.Next
		// 快指针向前移动2
		fast = fast.Next.Next
	}
	// 退出循环则快指针等于慢指针 存在闭环
	return true
}

func TestHasCycle(t *testing.T) {
	listNode1 := &ListNode{
		Val: 1,
	}
	listNode2 := &ListNode{
		Val: 2,
	}
	listNode3 := &ListNode{
		Val: 3,
	}
	listNode1.Next = listNode2
	listNode2.Next = listNode3
	listNode3.Next = listNode2
	listNode4 := &ListNode{Val: 4}
	listNode5 := &ListNode{Val: 5}
	listNode6 := &ListNode{Val: 6}
	listNode4.Next = listNode5
	listNode5.Next = listNode6
	fmt.Println(hasCycle(listNode1))
	fmt.Println(hasCycle(listNode4))
}
