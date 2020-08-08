package medium

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。



示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1. 排序后遍历记录每个数出现次数 时间复杂度O(nlogn) 空间复杂度O(1)
// 解法2. hash 时间复杂度O(1) 空间复杂度O(n)
func topKFrequent(nums []int, k int) []int {
	// 记录
	hashMap := make(map[int]int)
	for _, n := range nums {
		hashMap[n]++
	}

	// 置入小根堆
	h := &NodeHeap{}
	topK := min(k, len(hashMap))
	size := 0
	for k, v := range hashMap {
		if size < topK {
			heap.Push(h, &Node347{
				val:   k,
				times: v,
			})
			size++
		} else {
			if v > (*h)[0].times {
				heap.Pop(h)
				heap.Push(h, &Node347{
					val:   k,
					times: v,
				})
			}
		}
	}

	// get ans
	ans := make([]int, 0, topK)
	for i := 0; i < topK; i++ {
		ans = append(ans, heap.Pop(h).(*Node347).val)
	}
	return ans
}

type Node347 struct {
	val   int
	times int
}

// 实现内置库堆排序接口
type NodeHeap []*Node347

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].times < h[j].times
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node347))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestTopKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
