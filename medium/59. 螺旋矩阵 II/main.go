package main

import "fmt"

func generateMatrix(n int) [][]int {
	// 初始化结果二维数组
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	left, right, top, down := 0, n-1, 0, n-1
	index := 1
	for index <= n*n {
		if top <= down {
			for i := left; i <= right; i++ {
				res[top][i] = index
				index++
			}
			top++
		}

		if right >= left {
			for i := top; i <= down; i++ {
				res[i][right] = index
				index++
			}
			right--
		}

		if down >= top {
			for i := right; i >= left; i-- {
				res[down][i] = index
				index++
			}
			down--
		}

		if left <= right {
			for i := down; i >= top; i-- {
				res[i][left] = index
				index++
			}
			left++
		}
	}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
