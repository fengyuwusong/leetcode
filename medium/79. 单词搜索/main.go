package main

import "fmt"

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"SEE"
//输出：true
//
//
// 示例 3：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCB"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
// Related Topics数组 | 回溯 | 矩阵
//
// 👍 1377, 👎 0
//
//
//
//

// 2022-07-24 17:09:37 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
type Point struct {
	x int
	y int
}

// 定义上下左右四个点
var points = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var check func(i, j, currIndex int) bool
	check = func(i, j, currIndex int) bool {
		// 判断当前坐标是否相等
		if board[i][j] != word[currIndex] {
			return false
		}
		// 判断是否达到了最后一个字符
		if currIndex == len(word)-1 {
			return true
		}
		// 当前坐标设置为已访问
		visited[i][j] = true
		// 函数结束后 当前坐标设置为未访问
		defer func() {
			visited[i][j] = false
		}()
		// 步入下一个坐标
		for _, point := range points {
			// 判断是否越界
			if i+point.x < 0 || i+point.x >= m || j+point.y < 0 || j+point.y >= n {
				continue
			}
			// 判断是否已经访问过
			if visited[i+point.x][j+point.y] {
				continue
			}
			if check(i+point.x, j+point.y, currIndex+1) {
				return true
			}
		}
		return false
	}

	// 遍历所有的坐标起点
	for i := range board {
		for j := range board[i] {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))                                                        // true
	fmt.Println(exist(board, "SEE"))                                                       // true
	fmt.Println(exist(board, "ABCB"))                                                      // false
	fmt.Println(exist([][]byte{{'a', 'a'}}, "aaa"))                                        // false
	fmt.Println(exist([][]byte{{'c', 'a', 'a'}, {'a', 'a', 'a'}, {'b', 'c', 'd'}}, "aab")) // false
}
