package main

import "fmt"

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// Related Topics字符串 | 动态规划
//
// 👍 5583, 👎 0
//
//
//
//

// 2022-08-17 01:16:29 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	var res string
	for i := range s {
		// 判断是否是回文
		temp1 := isPalindrome(s, i, i)
		temp2 := isPalindrome(s, i, i+1)
		if len(temp1) > len(res) {
			res = temp1
		}
		if len(temp2) > len(res) {
			res = temp2
		}
	}
	return res
}

// isPalindrome 以 [l, r]为中心扩散返回最长的回文串
func isPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
