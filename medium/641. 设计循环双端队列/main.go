package main

import "fmt"

//设计实现双端队列。
//
// 实现 MyCircularDeque 类:
//
//
// MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
// boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
// boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
// boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
// boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
// int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
// int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
// boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false 。
// boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
//
//
//
//
// 示例 1：
//
//
//输入
//["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront",
//"getRear", "isFull", "deleteLast", "insertFront", "getFront"]
//[[3], [1], [2], [3], [4], [], [], [], [4], []]
//输出
//[null, true, true, true, false, 2, true, true, true, 4]
//
//解释
//MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= k <= 1000
// 0 <= value <= 1000
// insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty,
// isFull 调用次数不大于 2000 次
//
// Related Topics设计 | 队列 | 数组 | 链表
//
// 👍 127, 👎 0
//
//
//
//

// 2022-08-15 00:42:59 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
type MyCircularDeque struct {
	data []int
	len  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data: make([]int, 0, k),
		len:  k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.data)+1 > this.len {
		return false
	}
	this.data = append([]int{value}, this.data...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.data)+1 > this.len {
		return false
	}
	this.data = append(this.data, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[:len(this.data)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}

func (this *MyCircularDeque) GetRear() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.len == len(this.data)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	obj := Constructor(3)
	param_1 := obj.InsertLast(1)
	param_2 := obj.InsertLast(2)
	param_3 := obj.InsertFront(3)
	param_4 := obj.InsertFront(4)
	param_5 := obj.GetRear()
	param_6 := obj.IsFull()
	param_7 := obj.DeleteLast()
	param_8 := obj.InsertFront(4)
	param_9 := obj.GetFront()
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
	fmt.Println(param_7)
	fmt.Println(param_8)
	fmt.Println(param_9)

	o2 := Constructor(4)
	println(o2.InsertFront(9))
	println(o2.DeleteLast())
	println(o2.GetRear())
	println(o2.GetFront())
	println(o2.GetFront())
	println(o2.DeleteFront())
	println(o2.InsertFront(6))
	println(o2.InsertLast(5))
	println(o2.InsertFront(9))
	println(o2.GetFront())
	println(o2.InsertFront(6))
}
