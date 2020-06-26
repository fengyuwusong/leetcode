package easy

import "testing"

/**
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func strStr(haystack string, needle string) int {
	// 健壮性编写
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	var i, j int
	for i = 0; i < len(haystack); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
		// 剩余字符串长度小于needle长度则返回
		if len(haystack)-i-1 < len(needle) {
			return -1
		}
	}
	return -1
}

// kmp
func strStr2(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	next := buildNext(needle)
	i, j := 0, 0

	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func buildNext(s string) []int {
	// 初始化next数组
	next := make([]int, len(s))
	// 第一位赋值为-1 由于存在 j= next[j] 如第一位赋值0则会产生死循环
	next[0] = -1
	// 主串指针
	i := 0
	j := -1 // 模式串指针 赋值-1 与上面同理
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			// 模式串处于重开状态 或 当前字符相等 当前字符前进 赋值next表
			i++
			j++
			// 重开状态 next[i]=0 即主串当前字符并不是子串开头
			next[i] = j
		} else {
			// 否则模式串起始位置回归上一个最长前缀后缀相同位置
			j = next[j]
		}
	}
	return next
}

func TestStrStr(t *testing.T) {
	println(strStr("hello", "ll"))
	println(strStr("aaaaa", "bba"))
	println(strStr("aaaaa", "aa"))
	println(strStr("mississippi", "issipi"))

	println(strStr2("hello", "ll"))
	println(strStr2("aaaaa", "bba"))
	println(strStr2("aaaaa", "aa"))
	println(strStr2("mississippi", "issipi"))
	println(strStr2("mississippi", ""))
}
