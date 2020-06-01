package medium

import (
	"fmt"
	"math"
	"testing"
)

// bfs + 贪心
// 可将题目看为 给定n元树 求根节点到任一完全平方数节点的最小距离
func numSquares(n int) int {
	// 计算 [1, n]之间的所有完全平方数
	squareNumbers := make([]int, int(math.Sqrt(float64(n))))
	for i := 1; i < int(math.Sqrt(float64(n))+1); i++ {
		squareNumbers[i-1] = i * i
	}

	// 将queue定为set 可以避免相同节点的重复计算
	queueSet := make(map[int]struct{})
	queueSet[n] = struct{}{}
	// 记录当前层级
	level := 0

	// bfs
	for len(queueSet) != 0 {
		level++
		// 遍历queue 计算下一层级节点
		nextQueueSet := make(map[int]struct{})
		for node, _ := range queueSet {
			// 遍历所有完全平方数
			for _, v := range squareNumbers {
				// 找到对应节点为完全平方数
				if v == node {
					return level
				}
				// 大于节点值则跳过
				if v > node {
					continue
				}
				// 求剩余节点并加入队列中 队列中已有则不重复添加
				if _, ok := nextQueueSet[node-v]; !ok {
					nextQueueSet[node-v] = struct{}{}
				}
			}
		}
		queueSet = nextQueueSet
	}
	return level
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(12))
}
