package easy

import (
	"fmt"
	"testing"
)

/**
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。


示例:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false


说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 判断栈2是否为空 为空则将栈1元素转移到栈2
	if len(this.stack2) == 0 && len(this.stack1) != 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	// 获取栈2最后一个元素
	res := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	// 判断栈1是否为空 不为空则将元素转移到辅助栈
	if len(this.stack2) == 0 && len(this.stack1) != 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	return this.stack2[len(this.stack2)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	// 判断栈1是否为空 不为空则将元素转移到辅助栈
	if len(this.stack2) == 0 && len(this.stack1) != 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	return len(this.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func TestMyQueue(t *testing.T) {
	obj := ConstructorMyQueue()
	obj.Push(1)
	obj.Push(2)
	param_1 := obj.Peek()
	param_2 := obj.Pop()
	param_3 := obj.Empty()
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
}
