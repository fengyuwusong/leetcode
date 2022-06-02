package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return nil
	}
	m := len(matrix[0])

	// 定义初始化上下左右4个边界
	top, down, left, right := 0, n-1, 0, m-1
	var res []int

	for len(res) < n*m {
		// 如上边界小于等于下边界 则 左上 -> 右上
		if top <= down {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}

		// 如果左边界小于等于右边界 则 右上 -> 右下
		if left <= right {
			for i := top; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		// 如果下边界大于上边界 则 右下 -> 左下
		if down >= top {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down--
		}

		// 如果左边界小于等于右边界 则左下 -> 左上
		if left <= right {
			for i := down; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{6, 9, 7}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
