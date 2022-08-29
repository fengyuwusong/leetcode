package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type pair struct {
	sum, i int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less(i, j int) bool {
	return h[i].sum > h[j].sum
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func kSum(nums []int, k int) int64 {
	n := len(nums)
	// nums中所有非负数的和为sum 任一子序列的和都等价于从sum中减去某些非负数 / 加上某些负数得到
	sum := 0
	for i, num := range nums {
		if num >= 0 {
			sum += num
		} else {
			// 负数取绝对值 则可以统一成从sum中减去某些数得到任一子序列
			nums[i] = -num
		}
	}
	// 排序数组
	sort.Ints(nums)
	// 使用最大堆实现第K大的数 第一大子序列和为Sum
	h := &hp{{sum, 0}}
	// 如k大于1 则进入数组遍历得到其他大的数
	for ; k > 1; k-- {
		// 获取当前最大的子序列和
		p := heap.Pop(h).(pair)
		if p.i < n {
			// 当前最大子序列和 - 数组p.i的值 新元素p.i + 1
			// 保留nums[p.i-1]
			heap.Push(h, pair{p.sum - nums[p.i], p.i + 1})
			// 当前p.i > 0
			if p.i > 0 {
				// 不保留 nums[p.i-1]，把之前减去的加回来
				heap.Push(h, pair{p.sum - nums[p.i] + nums[p.i-1], p.i + 1})
			}
		}
	}
	return int64((*h)[0].sum)
}

func main() {
	fmt.Println(kSum([]int{2, 4, -2}, 5))
	fmt.Println(kSum([]int{1, -2, 3, 4, -10, 12}, 16))
}
