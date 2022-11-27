package main

import "fmt"

//在一个 n x n 的矩阵 grid 中，除了在数组 mines 中给出的元素为 0，其他每个元素都为 1。mines[i] = [xi, yi]表示
//grid[xi][yi] == 0
//
// 返回 grid 中包含 1 的最大的 轴对齐 加号标志的阶数 。如果未找到加号标志，则返回 0 。
//
// 一个 k 阶由 1 组成的 “轴对称”加号标志 具有中心网格 grid[r][c] == 1 ，以及4个从中心向上、向下、向左、向右延伸，长度为 k-1，
//由 1 组成的臂。注意，只有加号标志的所有网格要求为 1 ，别的网格可能为 0 也可能为 1 。
//
//
//
// 示例 1：
//
//
//
//
//输入: n = 5, mines = [[4, 2]]
//输出: 2
//解释: 在上面的网格中，最大加号标志的阶只能是2。一个标志已在图中标出。
//
//
// 示例 2：
//
//
//
//
//输入: n = 1, mines = [[0, 0]]
//输出: 0
//解释: 没有加号标志，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 500
// 1 <= mines.length <= 5000
// 0 <= xi, yi < n
// 每一对 (xi, yi) 都 不重复
//
// Related Topics数组 | 动态规划
//
// 👍 200, 👎 0
//
//
//
//

// fengyuwusong 2022-11-27 18:14:32
//leetcode submit region begin(Prohibit modification and deletion)
func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}
	banned := map[int]bool{}
	for _, p := range mines {
		banned[p[0]*n+p[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		/* left */
		for j := 0; j < n; j++ {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
		count = 0
		/* right */
		for j := n - 1; j >= 0; j-- {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		/* up */
		for j := 0; j < n; j++ {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
		}
		count = 0
		/* down */
		for j := n - 1; j >= 0; j-- {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
			ans = max(ans, dp[j][i])
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
}
