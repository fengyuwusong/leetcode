package medium

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/**
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。



示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的连续子数组。


进阶：

如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 解法1 暴力解 时间复杂度O(n^2) 空间复杂度O(1)
// 解法2 双指针 时间复杂度O(n) 空间复杂度O(1)
// 解法3 前缀和+二分查找 时间复杂度O(nlogn) 空间复杂度O(n)

// 解法2
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, i, j := 0, 0, 0
	ans := math.MaxInt64
	for j < len(nums) {
		sum += nums[j]
		// 取最小
		for sum >= s {
			if j-i+1 < ans {
				ans = j - i + 1
			}
			sum -= nums[i]
			i++
		}
		j++
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 解法3
// 计算前缀和数组后 根据题意, 我们需要求
// sums[k] - sums[j] >= s => sums[j] + s <= sums[k]
// 那么由于sums是递增的 我们仅需求出 sums[j]+s的值 然后二分法求出k值(近似=sums[j]+s)即可
// 即我们可以通过遍历j前缀, 通过二分法确定在哪个k的位置前缀和必定大于s
// 例如 sums = [0, 2, 5, 7, 10, 15, 18, 20]
// s=9
// j=0 sums[0] + 9 = 9 bound=4 即从0到第4个元素之后大于等于s ans=4
// j=1 sums[1] + 2 = 11 bound=5 即从1到第5个元素之后大于等于s ans=5-1+1=4
// ...
func minSubArrayLen3(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt64
	sums := make([]int, n+1)
	// 计算前缀和
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	// 遍历前缀和
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		// 二分查找如查找失败 则返回假如添加目标数后的位置
		// 例如 [2, 5, 7, 10, 15, 18, 20]
		// target=9 那么bound = 3
		bound := sort.SearchInts(sums, target)
		if bound <= n && ans > bound-i+1 {
			ans = bound - i + 1
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
func TestMinSubArrayLen(t *testing.T) {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
	fmt.Println(minSubArrayLen3(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen3(15, []int{1, 2, 3, 4, 5}))
}
