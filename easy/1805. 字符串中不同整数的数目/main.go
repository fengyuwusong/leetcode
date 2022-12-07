package main

import (
	"fmt"
	"unicode"
)

//给你一个字符串 word ，该字符串由数字和小写英文字母组成。
//
// 请你用空格替换每个不是数字的字符。例如，"a123bc34d8ef34" 将会变成 " 123 34 8 34" 。注意，剩下的这些整数为（相邻彼此至少有
//一个空格隔开）："123"、"34"、"8" 和 "34" 。
//
// 返回对 word 完成替换后形成的 不同 整数的数目。
//
// 只有当两个整数的 不含前导零 的十进制表示不同， 才认为这两个整数也不同。
//
//
//
// 示例 1：
//
//
//输入：word = "a123bc34d8ef34"
//输出：3
//解释：不同的整数有 "123"、"34" 和 "8" 。注意，"34" 只计数一次。
//
//
// 示例 2：
//
//
//输入：word = "leet1234code234"
//输出：2
//
//
// 示例 3：
//
//
//输入：word = "a1b01c001"
//输出：1
//解释："1"、"01" 和 "001" 视为同一个整数的十进制表示，因为在比较十进制值时会忽略前导零的存在。
//
//
//
//
// 提示：
//
//
// 1 <= word.length <= 1000
// word 由数字和小写英文字母组成
//
// Related Topics哈希表 | 字符串
//
// 👍 30, 👎 0
//
//
//
//

// fengyuwusong 2022-12-06 00:16:30
//leetcode submit region begin(Prohibit modification and deletion)
func numDifferentIntegers(word string) int {
	// 双指针
	memo := make(map[string]struct{})
	var l, r int
	for {
		// 左指针遍历到第一个数值字符
		for ; l < len(word) && !unicode.IsDigit(rune(word[l])); l++ {
		}
		if l == len(word) {
			break
		}
		// 右指针从l开始遍历到第一个非数值字符
		for r = l + 1; r < len(word) && unicode.IsDigit(rune(word[r])); r++ {
		}
		// 左指针去除0
		for r-l > 1 && word[l] == '0' {
			l++
		}
		// 记表
		memo[word[l:r]] = struct{}{}
		l = r
	}
	return len(memo)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(numDifferentIntegers("a1b01c001"))
	fmt.Println(numDifferentIntegers("leet1234code234"))
}
