package medium

import (
	"fmt"
	"testing"
)

/**
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？



示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 注意点：需不适用额外空间
// 方法一: 等式替换
// 		matrix[row][col] = matrix[col][n-row-1] n => len(matrix[0])
// =>   row = col | col = n-row-1
// =>   matrix[col][n-row-1] = matrix[n-row-1][n-col-1]
// =>   row = n-row-1 | col = n-col-1
// =>   matrix[n-row-1][n-col-1] = matrix[n-col-1][n-(n-row-1)-1] =  matrix[n-col-1][row]
// =>   row = n-col-1 | col = row
// =>   matrix[n-col-1][row] = matrix[row][n-(n-col-1)-1] = matrix[row][col]
// 发现最终等式相等 那么由于建立了四个等式 即替换了四次 最终仅需遍历四分之一的矩阵即可
// 综上所述，我们只需要枚举矩阵左上角高为 [n/2] 宽 [(n+1)/2]的子矩阵即可

// 方法二: 使用翻转代替旋转
// 假设矩阵 1 2 3
//         4 5 6
//         7 8 9
// => 以中间进行水平翻转
//         7 8 9
//         4 5 6
//         1 2 3
// => 以主对角线进行翻转
//		   7 4 1
//         8 5 2
//         9 6 3
// 两次翻转实际与方法一等式对应
func rotate(matrix [][]int) {
	matrixLen := len(matrix)
	if matrixLen == 0 {
		return
	}
	// 以中间进行水平翻转
	for r := 0; r < matrixLen/2; r++ {
		for c := 0; c < matrixLen; c++ {
			matrix[r][c], matrix[matrixLen-r-1][c] = matrix[matrixLen-r-1][c], matrix[r][c]
		}
	}
	// 以主对角线进行翻转
	for r := 0; r < matrixLen; r++ {
		for c := 0; c < r; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
}

func TestRotate(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
