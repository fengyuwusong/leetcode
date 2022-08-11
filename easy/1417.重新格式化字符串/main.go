package main

import "fmt"

// author: fengyuwusong date: 2022-08-11 15:04:15

// 1417. 重新格式化字符串
//给你一个混合了数字和字母的字符串 s，其中的字母均为小写英文字母。
//
// 请你将该字符串重新格式化，使得任意两个相邻字符的类型都不同。也就是说，字母后面应该跟着数字，而数字后面应该跟着字母。
//
// 请你返回 重新格式化后 的字符串；如果无法按要求重新格式化，则返回一个 空字符串 。
//
//
//
// 示例 1：
//
// 输入：s = "a0b1c2"
//输出："0a1b2c"
//解释："0a1b2c" 中任意两个相邻字符的类型都不同。 "a0b1c2", "0a1b2c", "0c2a1b" 也是满足题目要求的答案。
//
//
// 示例 2：
//
// 输入：s = "leetcode"
//输出：""
//解释："leetcode" 中只有字母，所以无法满足重新格式化的条件。
//
//
// 示例 3：
//
// 输入：s = "1229857369"
//输出：""
//解释："1229857369" 中只有数字，所以无法满足重新格式化的条件。
//
//
// 示例 4：
//
// 输入：s = "covid2019"
//输出："c2o0v1i9d"
//
//
// 示例 5：
//
// 输入：s = "ab123"
//输出："1a2b3"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 仅由小写英文字母和/或数字组成。
//
// Related Topics字符串
//
// 👍 57, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func reformat(s string) string {
	var numbersNum int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numbersNum++
		}
	}
	digitsNum := len(s) - numbersNum
	// 差值大于1 则返回空字符串
	if digitsNum-numbersNum > 1 || digitsNum-numbersNum < -1 {
		return ""
	}
	// 决定奇数是否存放数字
	ans := make([]byte, len(s))
	var flag bool
	if numbersNum > digitsNum {
		flag = true
	}
	// flag 则i填充数字，否则填充字母
	i, j := 0, 1
	for curr := 0; curr < len(s); curr++ {
		if s[curr] >= '0' && s[curr] <= '9' {
			if flag {
				ans[i] = s[curr]
				i += 2
			} else {
				ans[j] = s[curr]
				j += 2
			}
		} else {
			if flag {
				ans[j] = s[curr]
				j += 2
			} else {
				ans[i] = s[curr]
				i += 2
			}
		}

	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(reformat("a0b1c2"))
	fmt.Println(reformat("leetcode"))
	fmt.Println(reformat("1229857369"))
	fmt.Println(reformat("covid2019"))
	fmt.Println(reformat("ab123"))
	fmt.Println(reformat("1abc123"))
}
