package main

import "math"

// author: fengyuwusong date: 2022-08-24 15:51:15

// 456. 132 模式
//给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足
//：i < j < k 和 nums[i] < nums[k] < nums[j] 。
//
// 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4]
//输出：false
//解释：序列中不存在 132 模式的子序列。
//
//
// 示例 2：
//
//
//输入：nums = [3,1,4,2]
//输出：true
//解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
//
//
// 示例 3：
//
//
//输入：nums = [-1,3,2,0]
//输出：true
//解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 2 * 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
// Related Topics栈 | 数组 | 二分查找 | 有序集合 | 单调栈
//
// 👍 690, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func find132pattern(nums []int) bool {
	n := len(nums)
	// 单调递减栈
	var stack []int
	// 132 => ijk 模式中的k的值 初始化为最小
	k := math.MinInt64
	// 从后往前遍历
	for i := n - 1; i >= 0; i-- {
		// 如果当前元素小于k 这说明找到i了
		if nums[i] < k {
			return true
		}
		// 否则维护递减单调栈 单调栈不为空且当前元素大于栈顶元素 则出栈并更新k值
		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			k = max(k, stack[len(stack)-1]) // 说明栈顶元素能作为k值(比当前元素小) 更新k的最大值
			stack = stack[:len(stack)-1]    // 将可能的k值出栈
		}
		// 将当前元素加入栈 该元素为栈最小元素
		stack = append(stack, nums[i])
	}
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	println(find132pattern([]int{1, 2, 3}))           // false
	println(find132pattern([]int{-1, 3, 2, 0}))       // true
	println(find132pattern([]int{-1, -2, -3, 0}))     // false
	println(find132pattern([]int{-1, -2, -3, 0, -2})) // true
	println(find132pattern([]int{1, 0, 1, -4, -3}))   // false
	println(find132pattern([]int{-5, 0, 1, -4, -3}))  // true
	println(find132pattern([]int{3, 5, 0, 3, 4}))     // true
}
