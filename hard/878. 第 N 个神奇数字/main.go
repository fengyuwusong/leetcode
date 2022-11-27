package main

import (
	"fmt"
)

//一个正整数如果能被 a 或 b 整除，那么它是神奇的。
//
// 给定三个整数 n , a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案 对 10⁹ + 7 取模 后的值。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 1, a = 2, b = 3
//输出：2
//
//
// 示例 2：
//
//
//输入：n = 4, a = 2, b = 3
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
// 2 <= a, b <= 4 * 10⁴
//
//
//
// Related Topics数学 | 二分查找
//
// 👍 110, 👎 0
//
//
//
//

// fengyuwusong 2022-11-22 01:09:40
//leetcode submit region begin(Prohibit modification and deletion)
const mod int = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	l := min(a, b)
	r := n * min(a, b)
	c := a / gcd(a, b) * b
	for l <= r {
		mid := (l + r) >> 1
		cnt := mid/a + mid/b - mid/c
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(nthMagicalNumber(1, 2, 3))
	fmt.Println(nthMagicalNumber(4, 2, 3))
}
