package main

import "fmt"

// author: fengyuwusong date: 2022-08-03 15:43:45

// 73. 矩阵置零
//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
//
//
// 示例 2：
//
//
//输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
//输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -2³¹ <= matrix[i][j] <= 2³¹ - 1
//
//
//
//
// 进阶：
//
//
// 一个直观的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个仅使用常量空间的解决方案吗？
//
// Related Topics数组 | 哈希表 | 矩阵
//
// 👍 752, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	// 使用变量记录为零的行列
	m := len(matrix)
	n := len(matrix[0])
	var row, col []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}
	for i := 0; i < len(row); i++ {
		matrix[row[i]] = make([]int, n)
	}
	for i := 0; i < len(col); i++ {
		for j := 0; j < m; j++ {
			matrix[j][col[i]] = 0
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
