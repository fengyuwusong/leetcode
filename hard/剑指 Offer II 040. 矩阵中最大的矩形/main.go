package main

import "fmt"

//给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
//
// 注意：此题 matrix 输入格式为一维 01 字符串数组。
//
//
//
// 示例 1：
//
//
//
//
//输入：matrix = ["10100","10111","11111","10010"]
//输出：6
//解释：最大矩形如上图所示。
//
//
// 示例 2：
//
//
//输入：matrix = []
//输出：0
//
//
// 示例 3：
//
//
//输入：matrix = ["0"]
//输出：0
//
//
// 示例 4：
//
//
//输入：matrix = ["1"]
//输出：1
//
//
// 示例 5：
//
//
//输入：matrix = ["00"]
//输出：0
//
//
//
//
// 提示：
//
//
// rows == matrix.length
// cols == matrix[0].length
// 0 <= row, cols <= 200
// matrix[i][j] 为 '0' 或 '1'
//
//
//
//
// 注意：本题与主站 85 题相同（输入参数格式不同）： https://leetcode-cn.com/problems/maximal-
//rectangle/
// Related Topics栈 | 数组 | 动态规划 | 矩阵 | 单调栈
//
// 👍 52, 👎 0
//
//
//
//

// 2022-08-07 00:46:53 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 {
		return 0
	}
	// 生成每列的高度数组
	m := len(matrix)
	n := len(matrix[0])
	height := make([]int, n)
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// 计算最大矩形面积
		temp := calMaximalRectangle(height)
		if ans < temp {
			ans = temp
		}
	}
	return ans

}

func calMaximalRectangle(height []int) int {
	lIndex := make([]int, len(height))
	rIndex := make([]int, len(height))
	var stack []int
	for i := 0; i < len(height); i++ {
		// 如果栈不为空 且栈顶元素大于当前元素 则出栈 同时当前出栈元素的右边界为i
		for len(stack) > 0 && height[stack[len(stack)-1]] > height[i] {
			rIndex[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		// 栈为空 则当前元素左边界为-1
		if len(stack) == 0 {
			lIndex[i] = -1
		} else {
			// 否则当前元素的左边界为栈顶元素
			lIndex[i] = stack[len(stack)-1]
		}
		// 入栈
		stack = append(stack, i)
	}
	// 栈剩余元素的右边界为最后一个元素
	for len(stack) != 0 {
		rIndex[stack[len(stack)-1]] = len(height)
		stack = stack[:len(stack)-1]
	}
	// 计算最大矩形面积
	var ans int
	for i := 0; i < len(height); i++ {
		area := (rIndex[i] - lIndex[i] - 1) * height[i]
		if ans < area {
			ans = area
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maximalRectangle([]string{
		"10100",
		"10111",
		"11111",
		"10010",
	}))
}
