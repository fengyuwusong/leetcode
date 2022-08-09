package main

import (
	"fmt"
	"strings"
)

// author: fengyuwusong date: 2022-08-04 14:57:08

// 125. 验证回文串
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
// Related Topics双指针 | 字符串
//
// 👍 555, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < len(s) && !checkDigitAndWord(s[i]) {
			i++
		}
		for j >= 0 && !checkDigitAndWord(s[j]) {
			j--
		}
		if i < j && s[i] != s[j] {
			return false
		}
	}
	return true
}

func checkDigitAndWord(r byte) bool {
	if !(r >= 'a' && r <= 'z' || r >= '0' && r <= '9') {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(".,"))
}
