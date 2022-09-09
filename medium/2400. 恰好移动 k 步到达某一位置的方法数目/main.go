package main

import (
	"fmt"
)

type pair struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func numberOfWays(startPos int, endPos int, k int) int {
	dp := map[pair]int{}
	// 定义f(x, left) 表示当前在x 还剩left步时，走到终点的方案
	var f func(int, int) int
	f = func(x int, left int) int {
		if abs(x-endPos) > left {
			return 0
		}
		if left == 0 {
			return 1 // 成功到达终点 1种方案
		}
		p := pair{x, left}
		if v, ok := dp[p]; ok {
			return v
		}
		res := (f(x-1, left-1) + f(x+1, left-1)) % (1e9 + 7)
		dp[p] = res
		return res
	}
	return f(startPos, k)
}

func main() {
	fmt.Println(numberOfWays(272, 270, 6))
	fmt.Println(numberOfWays(1, 2, 3))
	fmt.Println(numberOfWays(1, 2, 5))
	fmt.Println(numberOfWays(2, 5, 10))
	fmt.Println(numberOfWays(1, 2, 999))
}
