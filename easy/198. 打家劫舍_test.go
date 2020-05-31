package easy

import (
	"fmt"
	"testing"
)

/**
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 动态规划 据题意可知
// 当遍历到i(i>2)时, dp[i] = max(dp[i-1], dp[i-2] + nums[i])
// 边界条件
// 1. len=1: 返回nums[0]
// 2. len=2: 返回max(nums[0], nums[1])
// 由推导式可知 只需记录 dp[i-1] 和dp[i-2] 两个记录即可 时间复杂度O(n) 空间复杂度O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 前两个数 last1=dp[i-1] last2=dp[i-1] 分别求i=1 i=2 时dp的值
	// 注意: 此时在边界条件中 len=1 dp[0]=nums[0] len=2 dp[1]=max(nuns[0], nums[1])
	last1, last2 := max(nums[0], nums[1]), nums[0]
	for i := 2; i < len(nums); i++ {
		// last1 置位当前遍历 i last2 置为之前的last1
		last1, last2 = max(last1, last2+nums[i]), last1
	}
	return last1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{1, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 3, 1}))
	fmt.Println(rob([]int{1, 1, 1, 2}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
