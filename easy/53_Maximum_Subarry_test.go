package easy

import (
	"fmt"
	"testing"
)

/**
Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划 重点为推导式 dp[i] = max(nums[i], dp[i-1] + nums[i]) dp为在每个位置上的最大子串和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(nums[i], dp[i-1]+nums[i])
	}
	// 遍历dp数组查找遍历到那个节点中值最大
	max := dp[0]
	for _, d := range dp {
		max = maxInt(max, d)
	}
	return max
}

// 贪心法 实际为动态规划 使用max记录数组中最大值
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = maxInt(nums[i], nums[i]+sum)
		max = maxInt(sum, max)
	}
	return max
}

// 分治法
func maxSubArray3(nums []int) int {
	// 递归出口
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	// 递归处理左边最大值
	leftMax := maxSubArray3(nums[:numsLen/2])
	// 递归处理右边最大值
	rightMax := maxSubArray3(nums[numsLen/2:])

	// 使用贪心法处理中间最大值 分别从 中心向左右拓张查找最大值 符合分治思路
	// 中心往左最大值
	leftStart := numsLen/2 - 1
	lSum, lMax := nums[leftStart], nums[leftStart]
	for i := leftStart - 1; i >= 0; i-- {
		lSum = nums[i] + lSum // 此处应该直接相加 因为不能再从单前节点开始 中间往外拓展应是连续的
		lMax = maxInt(lSum, lMax)
	}

	// 中心往右最大值
	rightStart := numsLen / 2
	rSum, rMax := nums[rightStart], nums[rightStart]
	for i := rightStart + 1; i < len(nums); i++ {
		rSum = nums[i] + rSum
		rMax = maxInt(rSum, rMax)
	}
	return maxInt(maxInt(leftMax, rightMax), lMax+rMax)
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1}))

	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
