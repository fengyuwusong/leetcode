package main

import "fmt"

// 先对角线镜像对称二维数组 然后翻转每一行即为数组旋转90度的结果
func rotate(matrix [][]int) {
	n := len(matrix)
	// 对角线镜像对称二维数组
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 反转二维数组每一行
	for _, arr := range matrix {
		for j, _ := range arr {
			if j >= n-j-1 {
				break
			}
			arr[j], arr[n-j-1] = arr[n-j-1], arr[j]
		}
	}
}

func main() {
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	m2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	rotate(m1)
	rotate(m2)

	fmt.Println(m1)
	fmt.Println(m2)
}
