package linkedlist

import (
	"fmt"
	"math"
	"testing"
)

/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。
如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。
如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type LinkedNode struct {
	prev *LinkedNode
	next *LinkedNode
	val  int
}

type MyLinkedList struct {
	head *LinkedNode
	tail *LinkedNode
	len  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := new(LinkedNode)
	tail := new(LinkedNode)
	head.val = math.MinInt32
	head.next = tail
	tail.val = math.MaxInt32
	tail.prev = head
	return MyLinkedList{head, tail, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.len || index < 0 {
		return -1
	}
	cur := this.head.next
	ans := -1
	for i := index; i >= 0; i-- {
		if cur == nil {
			break
		}
		ans = cur.val
		cur = cur.next
	}
	return ans
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	next := this.head.next
	node := new(LinkedNode)
	node.val = val
	node.prev = this.head
	node.next = next

	this.head.next = node
	next.prev = node

	this.len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	pre := this.tail.prev

	node := new(LinkedNode)
	node.val = val
	node.prev = pre
	node.next = this.tail

	this.tail.prev = node
	pre.next = node

	this.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.len {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.len {
		this.AddAtTail(val)
		return
	}

	// 遍历到需要插入位置的当前节点
	curr := this.head
	for i := 0; i <= index; i++ {
		curr = curr.next
	}

	prev := curr.prev

	node := new(LinkedNode)
	node.val = val
	node.prev = prev
	node.next = curr

	prev.next = node
	curr.prev = node

	this.len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}

	// 遍历到需要删除的节点
	curr := this.head
	for i := 0; i <= index; i++ {
		curr = curr.next
	}

	prev := curr.prev
	next := curr.next

	curr.prev.next = next
	curr.next.prev = prev

	this.len--
}

func (this *MyLinkedList) print() {
	cur := this.head
	for i := 0; i < this.len; i++ {
		cur = cur.next
		fmt.Printf("%d=>", cur.val)
	}
	fmt.Print("\n")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//["","","","","","","","","","get","deleteAtIndex","deleteAtIndex"]
//[,,,,[5],[6],[4]]
func TestMyLinkedList(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	fmt.Println(obj.Get(5))
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)
	obj.print()
}
