package easy

import (
	"fmt"
	"testing"
)

/**
Merge two sorted linked lists and return it as a new queue.
The new queue should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for singly-linked queue.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 循环遍历合并即可
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 参数健壮性编写
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var mergeList *ListNode
	var currentNode *ListNode
	// 选取头结点 复用 l1 l2 地址
	if l1.Val < l2.Val {
		mergeList = l1
		l1 = l1.Next
	} else {
		mergeList = l2
		l2 = l2.Next
	}
	currentNode = mergeList

	for {
		// 退出循环条件
		if l1 == nil {
			currentNode.Next = l2
			break
		}
		if l2 == nil {
			currentNode.Next = l1
			break
		}
		// 比较大小
		if l1.Val < l2.Val {
			currentNode.Next = l1
			l1 = l1.Next
			currentNode = currentNode.Next
			continue
		}
		currentNode.Next = l2
		currentNode = currentNode.Next
		l2 = l2.Next
	}
	return mergeList
}

func TestMergeTwoLists(t *testing.T) {
	l5 := &ListNode{16, nil}
	l4 := &ListNode{8, l5}
	l3 := &ListNode{4, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	ll3 := &ListNode{6, nil}
	ll2 := &ListNode{3, ll3}
	ll1 := &ListNode{1, ll2}
	res := mergeTwoLists(l1, ll1)
	for {
		if res == nil {
			break
		}
		fmt.Printf("=>%v", res.Val)
		res = res.Next
	}
}
