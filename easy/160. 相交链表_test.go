package easy

import (
	"fmt"
	"testing"
)

/**
编写一个程序，找到两个单链表相交的起始节点。

如下面的两个链表：



在节点 c1 开始相交。



示例 1：



输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。


示例 2：



输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。


示例 3：



输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。


注意：

如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1 哈希表 时间复杂度O(n) 空间复杂度O(n)
// 解法2.1 双指针 分别两个指针i j从两链表头节点出发，则会有以下两种情况
// 1. 两链表长度相等 那么i j相等时则为链表得交点
// 2. 两链表长度不等 那么当其中一个到达链尾时记录长度为 x,记录另一链表得位置 所在位置为y 且该链表走到链尾距离为 z
// 那么 z 及为两链表开头得长度差值 则让较长得链表先走 x-y 然后两链表再再同时遍历，第一个相同节点则为交点

// 解法2.2 基于2.1优化 两个指针都遍历一遍A和B，就对齐了长度 然后一起遍历最后一遍 遇到相等就返回，最后都为nil也返回.
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func TestGetIntersectionNode(t *testing.T) {
	headA1 := &ListNode{Val: 1}
	headA2 := &ListNode{Val: 2}
	headA3 := &ListNode{Val: 3}
	headB1 := &ListNode{Val: 4}
	headB2 := &ListNode{Val: 5}
	headC1 := &ListNode{Val: 6}
	headC2 := &ListNode{Val: 7}
	headA1.Next = headA2
	headA2.Next = headA3
	headA3.Next = headC1
	headB1.Next = headB2
	headB2.Next = headC1
	headC1.Next = headC2
	fmt.Println(getIntersectionNode(headA1, headB1).Val)
}
