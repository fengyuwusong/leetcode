package main

import (
	"container/heap"
	"fmt"
)

//你正在玩一个单人游戏，面前放置着大小分别为 a、b 和 c 的 三堆 石子。
//
// 每回合你都要从两个 不同的非空堆 中取出一颗石子，并在得分上加 1 分。当存在 两个或更多 的空堆时，游戏停止。
//
// 给你三个整数 a 、b 和 c ，返回可以得到的 最大分数 。
//
// 示例 1：
//
//
//输入：a = 2, b = 4, c = 6
//输出：6
//解释：石子起始状态是 (2, 4, 6) ，最优的一组操作是：
//- 从第一和第三堆取，石子状态现在是 (1, 4, 5)
//- 从第一和第三堆取，石子状态现在是 (0, 4, 4)
//- 从第二和第三堆取，石子状态现在是 (0, 3, 3)
//- 从第二和第三堆取，石子状态现在是 (0, 2, 2)
//- 从第二和第三堆取，石子状态现在是 (0, 1, 1)
//- 从第二和第三堆取，石子状态现在是 (0, 0, 0)
//总分：6 分 。
//
//
// 示例 2：
//
//
//输入：a = 4, b = 4, c = 6
//输出：7
//解释：石子起始状态是 (4, 4, 6) ，最优的一组操作是：
//- 从第一和第二堆取，石子状态现在是 (3, 3, 6)
//- 从第一和第三堆取，石子状态现在是 (2, 3, 5)
//- 从第一和第三堆取，石子状态现在是 (1, 3, 4)
//- 从第一和第三堆取，石子状态现在是 (0, 3, 3)
//- 从第二和第三堆取，石子状态现在是 (0, 2, 2)
//- 从第二和第三堆取，石子状态现在是 (0, 1, 1)
//- 从第二和第三堆取，石子状态现在是 (0, 0, 0)
//总分：7 分 。
//
//
// 示例 3：
//
//
//输入：a = 1, b = 8, c = 8
//输出：8
//解释：最优的一组操作是连续从第二和第三堆取 8 回合，直到将它们取空。
//注意，由于第二和第三堆已经空了，游戏结束，不能继续从第一堆中取石子。
//
//
//
//
// 提示：
//
//
// 1 <= a, b, c <= 10⁵
//
//
// Related Topics 贪心 数学 堆（优先队列） 👍 33 👎 0

// author: fengyuwusong
// create time: 2022-12-21 00:19:55
// leetcode submit region begin(Prohibit modification and deletion)
// 优先队列模板
type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	h2 := *h
	result := h2[len(h2)-1]
	*h = h2[:len(h2)-1]
	return result
}

func maximumScore(a int, b int, c int) int {
	h := &hp{a, b, c}
	heap.Init(h)
	var ans int
	for {
		// 拿出两个石子
		num1 := heap.Pop(h).(int)
		num2 := heap.Pop(h).(int)
		if num1 == 0 || num2 == 0 {
			break
		}
		ans++
		// 减一后重新加回去
		heap.Push(h, num1-1)
		heap.Push(h, num2-1)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maximumScore(2, 4, 6))
	fmt.Println(maximumScore(4, 4, 6))
	fmt.Println(maximumScore(1, 8, 8))
}
