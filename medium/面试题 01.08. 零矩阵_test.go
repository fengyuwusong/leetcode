package medium

import (
	"fmt"
	"testing"
)

/**
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。



示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zero-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 使用两个一维数组记录需要清零的行和列
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	var zeroRows, zeroCols []int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zeroRows = append(zeroRows, i)
				zeroCols = append(zeroCols, j)
			}
		}
	}
	for _, row := range zeroRows {
		for col := range matrix[row] {
			matrix[row][col] = 0
		}
	}
	for _, col := range zeroCols {
		for row := range matrix {
			matrix[row][col] = 0
		}
	}
}

func TestSetZero(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
