package main

import "fmt"

// author: fengyuwusong date: 2022-07-22 14:48:30

// 64. 最小路径和
//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// Related Topics数组 | 动态规划 | 矩阵
//
// 👍 1298, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	// 初始化dp 并计算边界的路径和
	dp := make([][]int, len(grid))
	dp[0] = make([]int, len(grid[0]))
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 计算整个dp
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			// 分别比较上边和左边的路径和最小值 则当前路径等于其中最小值+当前值
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
