package queue

import (
	"fmt"
	"testing"
)

type MyCircularQueue struct {
	head int
	tail int
	size int
	data []int
}

// todo 可增加length字段去除余数运算 例如 enqueue时判断
/**
if l == this.length {
	return false
}

if this.tail == this.length-1 {
	this.tail = 0
} else {
	this.tail++
}
this.data[this.tail] = value
*/

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head: -1,
		tail: -1,
		size: k,
		data: make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	// 注意 此处空则提前将head指针置于0 即队列头为第一个数值
	if this.IsEmpty() {
		this.head = 0
	}
	// 由于前面已判断队列是否满的情况 故此处 tail指针后移 入队即可
	this.tail = (this.tail + 1) % this.size
	this.data[this.tail] = value
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	// data空时
	if this.IsEmpty() {
		return false
	}
	// 队列中仅剩一个元素 需考虑此特殊情况
	if this.tail == this.head {
		// head tail 指针同时置于 -1
		this.head = -1
		this.tail = -1
		return true
	}
	// 由于前面已排除空和单个元素 故剩余情况都有2个或以上元素在队列中 仅需前移head指针出队即可
	this.head = (this.head + 1) % this.size
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	// 由于前面已定义当仅剩1个元素出队的时候 head tail 指针置 -1 故仅需判断即可
	if this.head == -1 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	// 当head tail指针相邻时 说明队列已满 不可以采用此判断 (this.head-1)%this.size == this.tail 负数求余为负数
	if (this.tail+1)%this.size == this.head {
		return true
	}
	return false
}

func TestQueue(t *testing.T) {
	//Your MyCircularQueue object will be instantiated and called as such:
	obj := Constructor(3)
	fmt.Println(obj.EnQueue(1))
	fmt.Println(obj.EnQueue(2))
	fmt.Println(obj.EnQueue(3))
	fmt.Println(obj.EnQueue(4))
}
