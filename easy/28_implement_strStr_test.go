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

func TestStrStr(t *testing.T) {
	println(strStr("hello", "ll"))
	println(strStr("aaaaa", "bba"))
	println(strStr("aaaaa", "aa"))
	println(strStr("mississippi", "issipi"))
}
