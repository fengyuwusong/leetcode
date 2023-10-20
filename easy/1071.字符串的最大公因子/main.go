package main

import (
	"fmt"
	"strings"
)

// 对于字符串 s 和 t，只有在 s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
//
// 给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
//
//
//
// 示例 1：
//
//
// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"
//
//
// 示例 2：
//
//
// 输入：str1 = "ABABAB", str2 = "ABAB"
// 输出："AB"
//
//
// 示例 3：
//
//
// 输入：str1 = "LEET", str2 = "CODE"
// 输出：""
//
//
//
//
// 提示：
//
//
// 1 <= str1.length, str2.length <= 1000
// str1 和 str2 由大写英文字母组成
//
//
// Related Topics 数学 字符串 👍 364 👎 0

// author: fengyuwusong
// create time: 2023-10-21 00:10:02
// leetcode submit region begin(Prohibit modification and deletion)
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	var res string
	for i := 1; i <= len(str1); i++ {
		if str1[:i] == str2[:i] {
			if len(str1)%len(str1[:i]) != 0 {
				continue
			}
			if strings.Repeat(str1[:i], len(str1)/len(str1[:i])) == str1 &&
				strings.Repeat(str1[:i], len(str2)/len(str1[:i])) == str2 {
				res = str1[:i]
			}
		} else {
			return res
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(gcdOfStrings("abcabc", "abc"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}
