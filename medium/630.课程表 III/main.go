package main

import (
	"fmt"
	"sort"
)

// author: fengyuwusong date: 2022-08-12 15:57:22

// 630. 课程表 III
//这里有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，其中 courses[i] = [durationi,
//lastDayi] 表示第 i 门课将会 持续 上 durationi 天课，并且必须在不晚于 lastDayi 的时候完成。
//
// 你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。
//
// 返回你最多可以修读的课程数目。
//
//
//
// 示例 1：
//
//
//输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
//输出：3
//解释：
//这里一共有 4 门课程，但是你最多可以修 3 门：
//首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
//第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
//第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
//第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。
//
// 示例 2：
//
//
//输入：courses = [[1,2]]
//输出：1
//
//
// 示例 3：
//
//
//输入：courses = [[3,2],[4,3]]
//输出：0
//
//
//
//
// 提示:
//
//
// 1 <= courses.length <= 10⁴
// 1 <= durationi, lastDayi <= 10⁴
//
// Related Topics贪心 | 数组 | 堆（优先队列）
//
// 👍 353, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
type PriorityQueue struct {
	data []int
	size int
}

func buildPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{
		data: make([]int, cap+1),
	}
}

func (pq *PriorityQueue) Push(x int) {
	pq.size++
	pq.data[pq.size] = x
	pq.up(pq.size)
}

func (pq *PriorityQueue) Pop() int {
	// 取出第一个元素并对调
	min := pq.data[1]
	pq.data[1], pq.data[pq.size] = pq.data[pq.size], pq.data[1]
	// 删除最后一个元素
	pq.data[pq.size] = 0
	pq.size--
	// 将第一个元素下沉
	pq.down(1)
	return min
}

// up 将第i个元素上浮
func (pq *PriorityQueue) up(i int) {
	// 当前元素比父节点还大，则交换并上浮 继续与上层比较
	for i > 1 && pq.data[i] > pq.data[i/2] {
		pq.data[i], pq.data[i/2] = pq.data[i/2], pq.data[i]
		i /= 2
	}
}

// down 将第i个元素下沉
func (pq *PriorityQueue) down(i int) {
	for i < pq.size {
		// 左节点
		j := 2 * i
		// 右节点比左节点大 j++
		if j+1 <= pq.size && pq.data[j+1] > pq.data[j] {
			j++
		}
		// 当前元素比左右节点都大，则退出
		if pq.data[i] >= pq.data[j] {
			break
		}
		// 交换并下沉
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		i = j
	}
}

func (pq *PriorityQueue) Empty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue) Size() int {
	return pq.size
}

func scheduleCourse(courses [][]int) int {
	// 按照课程的持续时间排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	// 堆
	pq := buildPriorityQueue(len(courses))
	//优先队列中所有课程的总时间
	total := 0
	// 按照关闭时间顺序遍历课程
	for _, course := range courses {
		// 优先队列中的课程的总时间 + 当前课程的时长小于等于当前课程的关闭时间，则加入优先队列
		if t := course[0]; total+t <= course[1] {
			pq.Push(t)
			total += t
		} else if !pq.Empty() && t < pq.data[1] {
			//否则如果优先队列不为空 且 当前课程的时长小于优先队列中的第一个课程(最大)的时长，则将当前课程的时长加入优先队列
			total -= pq.Pop()
			pq.Push(t)
			total += t
		}
	}
	return pq.Size()
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}))
	fmt.Println(scheduleCourse([][]int{{1, 2}}))
	fmt.Println(scheduleCourse([][]int{{3, 2}, {4, 3}}))
	fmt.Println(scheduleCourse([][]int{{100, 2}, {32, 50}}))
}
