package medium

import (
	"fmt"
	"testing"
)

/**
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
注意:

数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1：dfs 暴力解 时间复杂度 O(2^N) 空间复杂度 O(N)
// dfs暴力解
func findTargetSumWays(nums []int, S int) int {
	// 深度遍历有多少种组合可以达到目标和
	var dfs func(nums []int, target int) int
	dfs = func(nums []int, target int) int {
		// 递归出口
		if len(nums) == 0 {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}
		return dfs(nums[1:], target-nums[0]) + dfs(nums[1:], target+nums[0])
	}
	return dfs(nums, S)
}

// 解法2：动态规划 背包问题
// 给你元素和目标，从元素中选取组合达到目标，这样的问题，都可以归类到背包问题里。
// 由于数组为非负整数 那么数组所有数字相加或相减 可以目标公式项的上下限 例如数组[1, 1, 1] 的范围为 [-3, 3]
// 那么数组中任意数加减都会在这个范围内, 我们可以通过这个方式来进行动态规划
// i 为当前遍历的元素 j为加减后的目标值 dp[i][j] 为可以得到目标值的组合方式，那么可以得到转移方程
// dp[i][j1] = dp[i-1][j - nums[i]] + dp[i-1][j + nums[i]]  (j1 = j - nums[i] = j + nums[i])
// 当前目标值的组合方式 = 上一个目标值减去当前元素 = 当前目标值的操作次数 + 上一个目标值加上当前元素 = 当前目标值的操作次数
// 例如 [1,1,1]
// i = 0  dp[0][-1] = 1 | dp[0][1] = 1
// i = 1  dp[1][-2] = dp[0][-1-1] = 1 | dp[1][0] = dp[0][-1+1] + dp[0][1-1] = 2 | dp[1][2] = dp[0][1+1] =1
// ...

// 由于动态规划需要从上一个状态推导出下一个状态 而上述状态转移公式依赖了当前j1值,那么我们需稍作替换将j1替换为j
// dp[i][j - nums[i]] += dp[i-1][j] 即我们可以通过 上一个状态目标和的次数计算出当前状态目标和的次数
// 同理 dp[i][j + nums[i]] += dp[i-1][j]
// 即将思维逆转 例如 i = 1时  dp[0][-1] = 1 那么我们可以推倒出 dp[0][-2] = 1 | dp[0][0] = 1

// 由于dp[i] 至于 dp[i-1] 有关 故仅需使用两个一维数组分别记录 上一个状态的不同目标和 j 的组合方式次数即可
// 时间复杂度 O(N*sum) 空间复杂度：O(sum)
func findTargetSumWays2(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}

	// 求数组公式上下界
	sum := 0
	for _, n := range nums {
		sum += n
	}
	// 如果S大于数组之和 返回0
	if S > sum {
		return 0
	}
	// 生成对应dp数组 sum*2+1 为数组和的上下界 [-sum, sum] 中间有 0 需+1
	dp := make([]int, sum*2+1)
	// 初始化dp 由于 初始的数组的和不会超过1000 故保证dp数组索引不为负数 需加上1000
	dp[nums[0]+sum] = 1
	// 注意此处使用加, 避免第一个元素为nums[0]=0的情况 即nums[0]+sum = -nums[0]+sum
	dp[-nums[0]+sum] += 1
	// 从1开始遍历数组
	for i := 1; i < len(nums); i++ {
		// 初始化next 记录当前不同目标和的组合方式
		next := make([]int, sum*2+1)
		// 遍历数组和范围
		for j := -1 * sum; j <= sum; j++ {
			// 当目标和存在组合方式时进行运算下一个dp的值
			if dp[j+sum] > 0 {
				next[j-nums[i]+sum] += dp[j+sum]
				next[j+nums[i]+sum] += dp[j+sum]
			}
		}
		// 遍历结束 记录上一个目标和组合次数状态
		dp = next
	}
	return dp[S+sum]
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays2([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
