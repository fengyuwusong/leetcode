package main

import "fmt"

//给你一个整数数组 nums 和一个整数 k 。 nums 仅包含 0 和 1 。每一次移动，你可以选择 相邻 两个数字并将它们交换。
//
// 请你返回使 nums 中包含 k 个 连续 1 的 最少 交换次数。
//
//
//
// 示例 1：
//
// 输入：nums = [1,0,0,1,0,1], k = 2
//输出：1
//解释：在第一次操作时，nums 可以变成 [1,0,0,0,1,1] 得到连续两个 1 。
//
//
// 示例 2：
//
// 输入：nums = [1,0,0,0,0,0,1,1], k = 3
//输出：5
//解释：通过 5 次操作，最左边的 1 可以移到右边直到 nums 变为 [0,0,0,0,0,1,1,1] 。
//
//
// 示例 3：
//
// 输入：nums = [1,1,0,1], k = 2
//输出：0
//解释：nums 已经有连续 2 个 1 了。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 要么是 0 ，要么是 1 。
// 1 <= k <= sum(nums)
//
//
// Related Topics 贪心 数组 前缀和 滑动窗口 👍 60 👎 0

// author: fengyuwusong
// create time: 2022-12-18 00:27:04
// leetcode submit region begin(Prohibit modification and deletion)
func minMoves(nums []int, k int) int {
	// 改变思路 计算将 0 移出数组的次数
	//将数组变更为 zero 数组
	zeroNums := generateZeros(nums)
	// 构造前缀数组
	pre := generatePreSum(zeroNums)

	// 使用滑动窗口遍历 zeroNums 分别计算满足 k 个连续 1 需要移动 0 的次数
	// 滑动窗口区间应该是[0, k-2(zero每一个元素代表两个 1 的间隔，故需要k-2)]
	cost, left, right := 0, 0, k-2
	// 计算第一个窗口的解
	for i := left; i <= right; i++ {
		// 窗口每个移动 0 的开销为到左右启动的最小值 cost = min(i-left+1, right-i+1)
		cost += zeroNums[i] * min(i+1, right-i+1)
	}

	// cost 的变化可以通过中值变化得出 中值无需变化
	// 中值左边的都减乘系数 1（需要移动的距离变小）
	// 中值右边的都加乘系数 1 (需要移动的距离变大)
	minCost := cost
	// 开始滑动窗口
	for left, right := 1, left+k-1; right < len(zeroNums); left, right = left+1, right+1 {
		mid := (left + right) / 2
		cost -= getRangeSum(left-1, mid-1, pre)
		cost += getRangeSum(mid+k%2, right, pre)
		minCost = min(minCost, cost)
	}

	return minCost
}

// generateZeros 计算每个 1 之间间隔的 0 数量生成新的数组
// origin: [0, 1, 1, 0, 1, 0, 0, 1] => [0,1,2]
func generateZeros(nums []int) []int {
	var result []int
	lastOne := -1
	for i, num := range nums {
		if num == 1 && lastOne == -1 {
			lastOne = i
			continue
		}
		if num == 1 && lastOne != -1 {
			result = append(result, i-lastOne-1)
			lastOne = i
		}
	}
	return result
}

// generatePrefix 构造前缀和 用于快速求出区间和
func generatePreSum(nums []int) []int {
	pre := make([]int, len(nums)+1)
	for i, num := range nums {
		pre[i+1] = pre[i] + num
	}
	return pre
}

// getRangeSum 获取区间和
func getRangeSum(start int, end int, pre []int) int {
	return pre[end+1] - pre[start]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minMoves([]int{1, 0, 0, 1, 0, 1}, 2))
	fmt.Println(minMoves([]int{1, 0, 0, 0, 0, 0, 1, 1}, 3))
	fmt.Println(minMoves([]int{1, 1, 0, 1}, 2))
}
