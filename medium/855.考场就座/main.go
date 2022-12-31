package main

import (
	"container/list"
)

//在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。
//
// 当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，
//那么学生就坐在 0 号座位上。)
//
// 返回 ExamRoom(int N) 类，它有两个公开的函数：其中，函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位
//置；函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在
//座位 p 上。
//
//
//
// 示例：
//
// 输入：["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[]
//,[4],[]]
//输出：[null,0,9,4,2,null,5]
//解释：
//ExamRoom(10) -> null
//seat() -> 0，没有人在考场里，那么学生坐在 0 号座位上。
//seat() -> 9，学生最后坐在 9 号座位上。
//seat() -> 4，学生最后坐在 4 号座位上。
//seat() -> 2，学生最后坐在 2 号座位上。
//leave(4) -> null
//seat() -> 5，学生最后坐在 5 号座位上。
//
//
//
//
// 提示：
//
//
// 1 <= N <= 10^9
// 在所有的测试样例中 ExamRoom.seat() 和 ExamRoom.leave() 最多被调用 10^4 次。
// 保证在调用 ExamRoom.leave(p) 时有学生正坐在座位 p 上。
//
//
// Related Topics 设计 有序集合 堆（优先队列） 👍 147 👎 0

// author: fengyuwusong
// create time: 2022-12-30 00:47:31
// leetcode submit region begin(Prohibit modification and deletion)
// 使用双向链表 当坐下来的时候，
type ExamRoom struct {
	seats *list.List // 座位链表 记录每个被坐的位置
	n     int        // 座位数
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		seats: list.New(),
		n:     n - 1,
	}
}

func (this *ExamRoom) Seat() int {
	// 当还没入座的时候 默认坐 0 位置
	if this.seats.Len() == 0 {
		this.seats.PushFront(0)
		return 0
	}
	// 查找两个节点之间的最大值 其中点就是距离任一个点距离最大的座位
	e := this.seats.Front()
	pre := e.Value.(int)
	max := pre
	// 记录需要插入的元素得到位置和值
	addVal := 0
	addFront := e
	// 开始遍历链表 获取之间最大的距离
	e = e.Next()
	for ; e != nil; e = e.Next() {
		// 计算中点的距离
		val := e.Value.(int)
		distant := (val - pre) / 2
		if distant > max {
			// 记录最大值 和需要插入的节点位置和值
			max = distant
			addFront = e
			addVal = pre + distant
		}
		// 重置前置节点
		pre = val
	}
	// 考虑插入的位置是链表末尾 则distant=n-pre
	distant := this.n - pre
	if distant > max {
		this.seats.PushBack(this.n)
		return this.n
	}
	// 插入遍历链表的位置
	this.seats.InsertBefore(addVal, addFront)
	return addVal
}

func (this *ExamRoom) Leave(p int) {
	// 遍历链表进行删除
	for curr := this.seats.Front(); curr.Next() != nil; curr = curr.Next() {
		val := curr.Value.(int)
		if val == p {
			this.seats.Remove(curr)
			return
		}
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	room := Constructor(10)
	println(room.Seat())
	println(room.Seat())
	println(room.Seat())
	println(room.Seat())
	room.Leave(4)
	println(room.Seat())
}
