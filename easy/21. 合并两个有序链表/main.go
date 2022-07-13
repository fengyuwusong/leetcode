package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 双指针遍历
	temp := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
			temp = temp.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
			temp = temp.Next
		}
	}
	if list1 != nil {
		temp.Next = list1
	} else {
		temp.Next = list2
	}
	return dummy.Next
}

func main() {
	ml := mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}})
	for ml != nil {
		println(ml.Val)
		ml = ml.Next
	}
	fmt.Println("==============")
	ml = mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, nil)
	for ml != nil {
		println(ml.Val)
		ml = ml.Next
	}
}
