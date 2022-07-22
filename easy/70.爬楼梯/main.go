package main

import "fmt"

// author: fengyuwusong date: 2022-07-22 15:24:54

// 70. 爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//
// 示例 2：
//
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
//
//
//
//
// 提示：
//
//
// 1 <= n <= 45
//
// Related Topics记忆化搜索 | 数学 | 动态规划
//
// 👍 2538, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	dp := make([]int, n+1)
	var fun func(n int) int
	fun = func(n int) int {
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		if n > 0 {
			if dp[n] == 0 {
				dp[n] = fun(n-1) + fun(n-2)
			}
			return dp[n]
		}
		return 0
	}

	return fun(n)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
