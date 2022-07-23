package main

import "fmt"

// 动态规划 dp[i][j] 表示从 (0,0) 到 (i,j) 的路径数 dp[i][j] = dp[i-1][j] + dp[i][j-1] 上边的路径 + 左边的路径
// 当i=0 或 j=0时，dp[i][j] = 1 边界只有一条路径
func uniquePaths(m int, n int) int {
	// 初始化dp dp[0][i] dp[i][0] 初始化为 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	// 生成dp数组
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	// 返回终点的路径数量
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 3))
}
