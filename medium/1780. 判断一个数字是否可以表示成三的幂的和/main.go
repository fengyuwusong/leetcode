package main

import "fmt"

//给你一个整数 n ，如果你可以将 n 表示成若干个不同的三的幂之和，请你返回 true ，否则请返回 false 。
//
// 对于一个整数 y ，如果存在整数 x 满足 y == 3ˣ ，我们称这个整数 y 是三的幂。
//
//
//
// 示例 1：
//
// 输入：n = 12
//输出：true
//解释：12 = 3¹ + 3²
//
//
// 示例 2：
//
// 输入：n = 91
//输出：true
//解释：91 = 3⁰ + 3² + 3⁴
//
//
// 示例 3：
//
// 输入：n = 21
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁷
//
// Related Topics数学
//
// 👍 39, 👎 0
//
//
//
//

// fengyuwusong 2022-12-09 00:40:00
//leetcode submit region begin(Prohibit modification and deletion)
func checkPowersOfThree(n int) bool {
	// 转为3进制 如果某一位为2则说明不是若干个3的幂等和
	for n != 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(checkPowersOfThree(12))
	fmt.Println(checkPowersOfThree(91))
	fmt.Println(checkPowersOfThree(21))
}
