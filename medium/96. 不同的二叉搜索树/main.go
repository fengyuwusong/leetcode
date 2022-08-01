package main

import "fmt"

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：5
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
// Related Topics树 | 二叉搜索树 | 数学 | 动态规划 | 二叉树
//
// 👍 1876, 👎 0
//
//
//
//

// 2022-08-02 01:01:44 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}
	var count func(lo, hi int) int
	count = func(lo, hi int) int {
		if lo > hi {
			return 1
		}
		// 查找备忘录
		if memo[lo][hi] != 0 {
			return memo[lo][hi]
		}
		var res int
		for mid := lo; mid <= hi; mid++ {
			left := count(lo, mid-1)
			right := count(mid+1, hi)
			res += left * right
		}
		// 记录备忘录
		memo[lo][hi] = res
		return res
	}
	return count(1, n)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(numTrees(3))
}
