package main

import "fmt"

//给你一个大小为 rows x cols 的矩阵 mat，其中 mat[i][j] 是 0 或 1，请返回 矩阵 mat 中特殊位置的数目 。
//
// 特殊位置 定义：如果 mat[i][j] == 1 并且第 i 行和第 j 列中的所有其他元素均为 0（行和列的下标均 从 0 开始 ），则位置 (i,
//j) 被称为特殊位置。
//
//
//
// 示例 1：
//
// 输入：mat = [[1,0,0],
//            [0,0,1],
//            [1,0,0]]
//输出：1
//解释：(1,2) 是一个特殊位置，因为 mat[1][2] == 1 且所处的行和列上所有其他元素都是 0
//
//
// 示例 2：
//
// 输入：mat = [[1,0,0],
//            [0,1,0],
//            [0,0,1]]
//输出：3
//解释：(0,0), (1,1) 和 (2,2) 都是特殊位置
//
//
// 示例 3：
//
// 输入：mat = [[0,0,0,1],
//            [1,0,0,0],
//            [0,1,1,0],
//            [0,0,0,0]]
//输出：2
//
//
// 示例 4：
//
// 输入：mat = [[0,0,0,0,0],
//            [1,0,0,0,0],
//            [0,1,0,0,0],
//            [0,0,1,0,0],
//            [0,0,0,1,1]]
//输出：3
//
//
//
//
// 提示：
//
//
// rows == mat.length
// cols == mat[i].length
// 1 <= rows, cols <= 100
// mat[i][j] 是 0 或 1
//
// Related Topics数组 | 矩阵
//
// 👍 79, 👎 0
//
//
//
//

// fengyuwusong 2022-09-04 21:27:06
//leetcode submit region begin(Prohibit modification and deletion)
func numSpecial(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	rowSum := make([]int, n)
	colSum := make([]int, m)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			rowSum[i] += mat[i][j]
			colSum[j] += mat[i][j]
		}
	}

	var ans int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 && rowSum[i]+colSum[j] == 2 {
				ans++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	mat := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	fmt.Println(numSpecial(mat))
}
