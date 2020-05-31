package easy

import (
	"fmt"
	"testing"
)

/**
680. 验证回文字符串 Ⅱ

给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 左右两种情况都要考虑 即需循环遍历左右两边
*/
func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindromeStr(s, i, j-1) || isPalindromeStr(s, i+1, j)
		}
	}
	return true
}

func isPalindromeStr(s string, startIndex, endIndex int) bool {
	for ; startIndex < endIndex; startIndex, endIndex = startIndex+1, endIndex-1 {
		if s[startIndex] != s[endIndex] {
			return false
		}
	}
	return true
}

func TestValidPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("acbca"))
	fmt.Println(validPalindrome("deeeee"))
	fmt.Println(validPalindrome("hbakab"))
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))
}
