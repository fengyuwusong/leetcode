package main

import "fmt"

// author: fengyuwusong date: 2022-08-10 10:19:57

// 37. 解数独
//编写一个程序，通过填充空格来解决数独问题。
//
// 数独的解法需 遵循如下规则：
//
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".
//",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".
//","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6
//"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[
//".",".",".",".","8",".",".","7","9"]]
//输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8
//"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],[
//"4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9",
//"6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4",
//"5","2","8","6","1","7","9"]]
//解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//
//
//
//
//
//
// 提示：
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字或者 '.'
// 题目数据 保证 输入数独仅有一个解
//
//
//
//
// Related Topics数组 | 回溯 | 矩阵
//
// 👍 1353, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	// 记录行 列 3*3 已存在的数字
	rowMemo := make([][]bool, 9)
	colMemo := make([][]bool, 9)
	matrixMemo := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowMemo[i] = make([]bool, 10)
		colMemo[i] = make([]bool, 10)
		matrixMemo[i] = make([]bool, 10)
	}
	// 初始化memo
	for i := range board {
		for j := range board {
			if board[i][j] != '.' {
				matrixNum := colMatrixNum(i, j)
				index := b2i(board[i][j])
				rowMemo[i][index] = true
				colMemo[j][index] = true
				matrixMemo[matrixNum][index] = true
			}
		}
	}
	// 定义回溯算法
	var backtrack func(row, col int) bool
	backtrack = func(row, col int) bool {
		// 所有空格均已填充
		if row == 9 && col == 9 {
			return true
		}
		// 歩进下一层
		if col == 9 {
			col = 0
			row++
		}
		for ; row < 9; row++ {
			for ; col < 9; col++ {
				if board[row][col] == '.' {
					matrixNum := colMatrixNum(row, col)
					// 从1-9开始选择
					for i := 1; i < 10; i++ {
						// 已存在
						if rowMemo[row][i] || colMemo[col][i] || matrixMemo[matrixNum][i] {
							continue
						}
						// 选择
						rowMemo[row][i] = true
						colMemo[col][i] = true
						matrixMemo[matrixNum][i] = true
						board[row][col] = byte('0' + i)
						if backtrack(row, col+1) {
							return true
						}
						// 撤回
						rowMemo[row][i] = false
						colMemo[col][i] = false
						matrixMemo[matrixNum][i] = false
						board[row][col] = '.'
					}
					// 如遍历完毕还是没选择 则回溯
					return false
				}
			}
			col = 0
		}
		// 如果遍历完毕 则说明矩阵全都填上了
		return true
	}
	backtrack(0, 0)
}

func b2i(v byte) int {
	return int(v - '0')
}

func colMatrixNum(row, col int) int {
	r := row / 3
	c := col / 3
	return r*3 + c
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}
