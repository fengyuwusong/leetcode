package main

import (
	"fmt"
	"unicode"
)

//给你一个混合字符串 s ，请你返回 s 中 第二大 的数字，如果不存在第二大的数字，请你返回 -1 。
//
// 混合字符串 由小写英文字母和数字组成。
//
//
//
// 示例 1：
//
//
//输入：s = "dfa12321afd"
//输出：2
//解释：出现在 s 中的数字包括 [1, 2, 3] 。第二大的数字是 2 。
//
//
// 示例 2：
//
//
//输入：s = "abc1111"
//输出：-1
//解释：出现在 s 中的数字只包含 [1] 。没有第二大的数字。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 只包含小写英文字母和（或）数字。
//
// Related Topics哈希表 | 字符串
//
// 👍 48, 👎 0
//
//
//
//

// fengyuwusong 2022-12-03 23:57:22
//leetcode submit region begin(Prohibit modification and deletion)
func secondHighest(s string) int {
	first, second := -1, -1
	for _, c := range s {
		if unicode.IsDigit(c) {
			num := int(c - '0')
			if num > first {
				second = first
				first = num
			} else if second < num && num < first {
				second = num
			}
		}
	}
	return second
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(secondHighest("a1v4232c0wqer421711"))
}
