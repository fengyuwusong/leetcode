package medium

import (
	"fmt"
	"math"
	"testing"
)

/**
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1:
输入:

0 0 0
0 1 0
0 0 0
输出:

0 0 0
0 1 0
0 0 0
示例 2:
输入:

0 0 0
0 1 0
1 1 1
输出:

0 0 0
0 1 0
1 2 1
注意:

给定矩阵的元素个数不超过 10000。
给定矩阵中至少有一个元素是 0。
矩阵中的元素只在四个方向上相邻: 上、下、左、右。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/01-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 使用bfs解决
// 使用bfs从0开始拓展每个非0节点的距离
// 需考虑多个0的情况 需将所有0看为一个整体 与矩阵中所有0相连 赋予「超级零」的概念
// 任意一个 1 到它最近的 0 的距离，会等于这个 1 到「超级零」的距离减去一。
// 由于只有一个超级零节点，这个「超级零」只和矩阵中的 0 相连，
// 所以在广度优先搜索的第一步中，「超级零」会被弹出队列，而所有的 0 会被加入队列，它们到「超级零」的距离为 1。
// 等价于一开始我们就将所有的 0 加入队列，它们的初始距离为 0。
// 这样以来，在广度优先搜索的过程中，我们每遇到一个 1，就得到了它到「超级零」的距离减去一，也就是 这个 1 到最近的 0 的距离。
// 时间复杂度 O(rc) 长宽乘积 空间复杂度O(1)
func updateMatrix(matrix [][]int) [][]int {
	row := len(matrix)
	if row == 0 {
		return nil
	}
	col := len(matrix[0])
	ans := make([][]int, row)
	visit := make([][]bool, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]bool, col)
		ans[i] = make([]int, col)
	}
	// 将所有的 0 添加进初始队列
	var queue []int
	for i, rv := range matrix {
		for j, cv := range rv {
			if cv == 0 {
				visit[i][j] = true
				queue = append(queue, i, j)
			}
		}
	}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	// 开始bfs
	for len(queue) != 0 {
		x, y := queue[0], queue[1]
		queue = queue[2:]
		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			// 不越界且未访问过
			if xx >= 0 && xx < row && yy >= 0 && yy < col && !visit[xx][yy] {
				ans[xx][yy] = ans[x][y] + 1
				queue = append(queue, xx, yy)
				visit[xx][yy] = true
			}
		}
	}
	return ans
}

// dp 动态规划解法
// 根据题意，我们可以取得0的位置，遍历二维数组，求出当前节点到达任一0的距离，然后取最短距离，但其复杂度会相当的高
// 为优化，我们可以从矩阵的四个角进行移动, 分别从左上、右上、左下、右下方向搜索
// 那么我们分别可以的到当前节点到达左上、右上、左下、右下方向0的最短距离
// 计算四个方向的最短距离 如此来我们就能确保当前到达的最近0距离最短
// 由此我们可以得到动态规划方程 dp[i] = min(dp[i-1] +1, dp[i+1] +1, dp[j-1]+1, dp[j+1]-1) (当前节点不等于0)
// i+1 i-1 j+1 j-1 分别对应当前节点的上下左右四个方向节点
// 当当前节点等于0时, dp[i]=0
// 优化点: 由于使用dp计算时，仅需计算右下和左上即可(任意对角)
// 原理：当计算左上时, 已经记录了任一节点到达左上节点的最短距离，那么计算右下时，当前节点的右下任一节点已包含了左上节点的最短距离，
// 实现了二维矩阵的全覆盖
func updateMatrix2(matrix [][]int) [][]int {
	row := len(matrix)
	if row == 0 {
		return nil
	}
	col := len(matrix[0])
	dp := make([][]int, row)
	// 初始化dp数组为无穷大
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < col; j++ {
			dp[i][j] = 10000
		}
	}
	// 从左上节点起始 遍历二维数组计算当前节点到达左上节点0的最短距离
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			}
			// 上节点无越界
			if i-1 >= 0 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j])+1))
			}
			// 左节点无越界
			if j-1 >= 0 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i][j-1])+1))
			}
		}
	}
	// 从右下节点起始 遍历二维数组计算当前节点到达右下节点0的最短距离
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			}
			// 下节点无越界
			if i+1 < row {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i+1][j])+1))
			}
			// 右节点无越界
			if j+1 < col {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i][j+1])+1))
			}
		}
	}
	return dp
}

func TestUpdateMatrix(*testing.T) {
	fmt.Println(updateMatrix([][]int{{0, 1, 1, 1}, {1, 1, 1, 0}, {0, 1, 1, 1}}))
	fmt.Println(updateMatrix2([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}
