package main

import "fmt"

//给你两个整数 n 和 k ，请你构造一个答案列表 answer ，该列表应当包含从 1 到 n 的 n 个不同正整数，并同时满足下述条件：
//
//
// 假设该列表是 answer = [a1, a2, a3, ... , an] ，那么列表 [|a1 - a2|, |a2 - a3|, |a3 - a4|
//, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数。
//
//
// 返回列表 answer 。如果存在多种答案，只需返回其中 任意一种 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, k = 1
//输出：[1, 2, 3]
//解释：[1, 2, 3] 包含 3 个范围在 1-3 的不同整数，并且 [1, 1] 中有且仅有 1 个不同整数：1
//
//
// 示例 2：
//
//
//输入：n = 3, k = 2
//输出：[1, 3, 2]
//解释：[1, 3, 2] 包含 3 个范围在 1-3 的不同整数，并且 [2, 1] 中有且仅有 2 个不同整数：1 和 2
//
//
//
//
// 提示：
//
//
// 1 <= k < n <= 10⁴
//
//
//
// Related Topics数组 | 数学
//
// 👍 118, 👎 0
//
//
//
//

// fengyuwusong 2022-09-08 01:05:24
//leetcode submit region begin(Prohibit modification and deletion)
// k=1 则数列为[1,2,3, ... ,n] 计算后数列=[1]
// k=n-1 [1,k+1,2,k,3,k-1 ...]  => [n-k, n, n-k+1, n-1, n-k+2 ...]
// 对于其它的一般情况，我们可以将这两种特殊情况进行合并，即列表的前半部分相邻差均为 1，后半部分相邻差从 k 开始逐渐递减到 1，这样从 1 到 k 的差值均出现一次，对应的列表即为：

func constructArray(n int, k int) []int {
	ans := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
}
