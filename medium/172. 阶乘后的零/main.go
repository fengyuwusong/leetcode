package main

import "fmt"

//给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
// 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：0
//解释：3! = 6 ，不含尾随 0
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：1
//解释：5! = 120 ，有一个尾随 0
//
//
// 示例 3：
//
//
//输入：n = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁴
//
//
//
//
// 进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
// Related Topics数学
//
// 👍 703, 👎 0
//
//
//
//

// fengyuwusong 2022-08-28 00:21:02
//leetcode submit region begin(Prohibit modification and deletion)
// 由于 2*5 = 10 故问题可变更为阶乘中分解因子 2和5 的最小个数，由于每个偶数都能分解为2的倍数，故问题变为求分解因子5的最小个数
func trailingZeroes(n int) int {
	curr := 5
	ans := 0
	for curr <= n {
		ans += n / curr
		curr *= 5
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
	fmt.Println(trailingZeroes(125))
}
