package main

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}

	// 构造前缀矩阵
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		preSum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			// 当前二维矩阵点的值等于其左边矩阵+上边矩阵-左上重复的矩阵+当前点的值
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}

func main() {
	nm := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	print(nm.SumRegion(2, 1, 4, 3))
}
