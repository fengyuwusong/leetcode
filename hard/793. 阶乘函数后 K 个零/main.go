package main

import (
	"fmt"
	"math"
)

//f(x) 是 x! 末尾是 0 的数量。回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 。
//
//
// 例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，因为 11!= 39916800 末端有 2 个 0 。
//
//
// 给定 k，找出返回能满足 f(x) = k 的非负整数 x 的数量。
//
//
//
// 示例 1：
//
//
//输入：k = 0
//输出：5
//解释：0!, 1!, 2!, 3!, 和 4! 均符合 k = 0 的条件。
//
//
// 示例 2：
//
//
//输入：k = 5
//输出：0
//解释：没有匹配到这样的 x!，符合 k = 5 的条件。
//
// 示例 3:
//
//
//输入: k = 3
//输出: 5
//
//
//
//
// 提示:
//
//
// 0 <= k <= 10⁹
//
// Related Topics数学 | 二分查找
//
// 👍 97, 👎 0
//
//
//
//

// fengyuwusong 2022-08-28 00:13:20
//leetcode submit region begin(Prohibit modification and deletion)
// x的阶乘0的个数取决于分解因子后5的个数 故问题变为分解因子5个数为k 这样的x有多少个
// tailZeroNums 计算n的阶乘后0的数量
func tailZeroNums(n int) int {
	ans := 0
	curr := 5
	for curr <= n {
		ans += n / curr
		curr *= 5
	}
	return ans
}

// preimageSizeFZFc 二分查找[0, int_max]中阶乘为k个0的数量(计算左右边界)
func preimageSizeFZF(k int) int {
	return rbound(k) - lbound(k) + 1
}

// 找左边界
func lbound(k int) int {
	l, r := 0, math.MaxInt
	var zeroNum int
	for l < r {
		mid := (r-l)/2 + l
		zeroNum = tailZeroNums(mid)
		if zeroNum == k {
			r = mid
		}
		if zeroNum < k {
			l = mid + 1
		}
		if zeroNum > k {
			r = mid
		}
	}
	return l
}

// 找右边界
func rbound(k int) int {
	l, r := 0, math.MaxInt
	for l < r {
		mid := (r-l)/2 + l
		zeroNum := tailZeroNums(mid)
		if zeroNum == k {
			l = mid + 1
		}
		if zeroNum < k {
			l = mid + 1
		}
		if zeroNum > k {
			r = mid
		}
	}
	return l - 1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(preimageSizeFZF(0))
	fmt.Println(preimageSizeFZF(5))
	fmt.Println(preimageSizeFZF(3))
}
