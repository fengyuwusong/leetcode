package main

import (
	"fmt"
)

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
// 示例 1:
//
//
//
//
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//
// 示例 2：
//
//
//
//
//输入： heights = [2,4]
//输出： 4
//
//
//
// 提示：
//
//
// 1 <= heights.length <=10⁵
// 0 <= heights[i] <= 10⁴
//
// Related Topics栈 | 数组 | 单调栈
//
// 👍 2088, 👎 0
//
//
//
//

// 2022-08-02 00:50:43 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
// 单调栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	// 记录从左到右边界
	var monoStack []int
	for i := 0; i < n; i++ {
		// 栈大于0且栈顶元素大于等于当前元素 出栈
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			// 弹出的栈顶元素的右边界等于当前遍历的元素位置
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		// 当前栈空 则左边界=-1
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			// 记录当前左边界
			left[i] = monoStack[len(monoStack)-1]
		}
		// 当前柱子位置入栈
		monoStack = append(monoStack, i)
	}
	// 栈剩下的位置右边界为n 不弹出说明可以拓展到最右边
	for len(monoStack) > 0 {
		right[monoStack[len(monoStack)-1]] = n
		monoStack = monoStack[:len(monoStack)-1]
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
	fmt.Println(largestRectangleArea([]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}))
}
