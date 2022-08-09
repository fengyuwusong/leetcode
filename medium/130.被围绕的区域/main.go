package main

import "fmt"

// author: fengyuwusong date: 2022-08-04 16:31:12

// 130. 被围绕的区域
//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
//
//
//
//
// 示例 1：
//
//
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
//"X","X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//
//
// 示例 2：
//
//
//输入：board = [["X"]]
//输出：[["X"]]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] 为 'X' 或 'O'
//
//
//
// Related Topics深度优先搜索 | 广度优先搜索 | 并查集 | 数组 | 矩阵
//
// 👍 837, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
// dfs 将边界的0置为*并扩展 然后再次遍历 剩余的0即是被围绕的区域 置为X
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 判断边界
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return
		}
		if board[i][j] == 'O' {
			board[i][j] = '*'
			// 四个方向
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(0, j)
		dfs(len(board)-1, j)
	}
	for i := 0; i < len(board); i++ {
		dfs(i, 0)
		dfs(i, len(board[0])-1)
	}
	// 将0置为X 将*置为O
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}
