package easy

import (
	"fmt"
	"testing"
)

/**
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-stack-using-queues
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type MyStack struct {
	list1 []int
	list2 []int
}

/** Initialize your data structure here. */
func ConstructorMyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	// 先将当前队列元素复制到list2
	this.list2 = this.list1
	// 将最新值先入list1
	this.list1 = []int{x}
	// 将list2值依次入队list1
	for len(this.list2) != 0 {
		this.list1 = append(this.list1, this.list2[0])
		this.list2 = this.list2[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.list1[0]
	this.list1 = this.list1[1:]
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.list1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.list1) == 0
}

func TestMyStack(t *testing.T) {
	stack := MyStack{}
	stack.Push(1)
	stack.Push(2)
	param1 := stack.Pop()
	param2 := stack.Top()
	param3 := stack.Empty()
	param4 := stack.Pop()
	param5 := stack.Empty()
	fmt.Println(param1)
	fmt.Println(param2)
	fmt.Println(param3)
	fmt.Println(param4)
	fmt.Println(param5)
}
