package main

import (
	"fmt"
	"math"
)

//给定数组 nums 和一个整数 k 。我们将给定的数组 nums 分成 最多 k 个相邻的非空子数组 。 分数 由每个子数组内的平均值的总和构成。
//
// 注意我们必须使用 nums 数组中的每一个数进行分组，并且分数不一定需要是整数。
//
// 返回我们所能得到的最大 分数 是多少。答案误差在 10⁻⁶ 内被视为是正确的。
//
//
//
// 示例 1:
//
//
//输入: nums = [9,1,2,3,9], k = 3
//输出: 20.00000
//解释:
//nums 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20.
//我们也可以把 nums 分成[9, 1], [2], [3, 9].
//这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.
//
//
// 示例 2:
//
//
//输入: nums = [1,2,3,4,5,6,7], k = 4
//输出: 20.50000
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
// Related Topics数组 | 动态规划 | 前缀和
//
// 👍 351, 👎 0
//
//
//
//

// fengyuwusong 2022-11-28 23:22:14
//leetcode submit region begin(Prohibit modification and deletion)
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	// 计算前缀和
	prefix := make([]float64, n+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + float64(x)
	}
	// 初始化dp dp[i][j] = 在区间[0, i)中 切分为j块的最大平均值和的分组
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	for i := 0; i <= n; i++ {
		// dp[i][1] 为前缀和对应的平均数
		dp[i][1] = prefix[i] / float64(i)
	}

	// dp[i][j] =  max(dp[i][j], dp[x][j-1] + 对应区间[x, i]的平均数) [x = 每个j开始的位置] 即取仅剩最后一个分快的dp获取最终值
	// 此处dp[i][j]仅使用j-1的数据，故可简化为1维数组
	for j := 2; j <= k; j++ {
		for i := j; i <= n; i++ {
			for x := j - 1; x < i; x++ {
				dp[i][j] = math.Max(dp[i][j], dp[x][j-1]+(prefix[i]-prefix[x])/float64(i-x))
			}
		}
	}
	return dp[n][k]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
}
