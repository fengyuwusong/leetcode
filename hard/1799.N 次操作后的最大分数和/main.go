package main

import (
	"fmt"
	"math/bits"
)

//给你 nums ，它是一个大小为 2 * n 的正整数数组。你必须对这个数组执行 n 次操作。
//
// 在第 i 次操作时（操作编号从 1 开始），你需要：
//
//
// 选择两个元素 x 和 y 。
// 获得分数 i * gcd(x, y) 。
// 将 x 和 y 从 nums 中删除。
//
//
// 请你返回 n 次操作后你能获得的分数和最大为多少。
//
// 函数 gcd(x, y) 是 x 和 y 的最大公约数。
//
//
//
// 示例 1：
//
// 输入：nums = [1,2]
//输出：1
//解释：最优操作是：
//(1 * gcd(1, 2)) = 1
//
//
// 示例 2：
//
// 输入：nums = [3,4,6,8]
//输出：11
//解释：最优操作是：
//(1 * gcd(3, 6)) + (2 * gcd(4, 8)) = 3 + 8 = 11
//
//
// 示例 3：
//
// 输入：nums = [1,2,3,4,5,6]
//输出：14
//解释：最优操作是：
//(1 * gcd(1, 5)) + (2 * gcd(2, 4)) + (3 * gcd(3, 6)) = 1 + 4 + 9 = 14
//
//
//
//
// 提示：
//
//
// 1 <= n <= 7
// nums.length == 2 * n
// 1 <= nums[i] <= 10⁶
//
//
// Related Topics 位运算 数组 数学 动态规划 回溯 状态压缩 数论 👍 34 👎 0

// author: fengyuwusong
// create time: 2022-12-22 00:21:28
// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(nums []int) int {
	n := len(nums)
	// f[s] 代表当数组处于 s不同位数已经被选择的时候 能取得的最大分数
	f := make([]int, 1<<n)
	// s 用于表示数组第 i 位是否已被选择删除 1 => 未被删除
	// 则其中 f[2^n-1] 为数组nums[]能取得的最大分数
	// 遍历f[s]的各种状态 2^n-1 dp计算最大值
	// f[next] = max(f[next], f[s]+c) => (当前分数+已选取的公约数)
	for s := 0; s < 1<<n; s++ {
		cnt := bits.OnesCount(uint(s))
		if cnt&1 == 1 {
			continue
		}
		for i := 0; i < n; i++ {
			if s>>i&1 == 1 {
				continue
			}
			for j := i + 1; j < n; j++ {
				if s>>j&1 == 1 {
					continue
				}
				// [i, j] 位置均为 0 则可选
				next := s | 1<<i | 1<<j
				// 当前可增加的分数 = 选取两个数的公约数 + 第 x => (cnt/2+1) 次选择
				// 此处 gcd 可用数组维护 避免重复计算
				c := gcd(nums[i], nums[j]) * (cnt/2 + 1)
				f[next] = max(f[next], f[s]+c)
			}
		}
	}
	return f[1<<n-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maxScore([]int{1, 2}))
	fmt.Println(maxScore([]int{3, 4, 6, 8}))
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6}))
}
