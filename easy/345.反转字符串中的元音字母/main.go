package main

import (
	"fmt"
)

// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
//
//
//
// 示例 1：
//
//
// 输入：s = "hello"
// 输出："holle"
//
//
// 示例 2：
//
//
// 输入：s = "leetcode"
// 输出："leotcede"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由 可打印的 ASCII 字符组成
//
//
// Related Topics 双指针 字符串 👍 333 👎 0

// author: fengyuwusong
// create time: 2023-10-24 00:33:20
// leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	bs := ([]byte)(s)
	dictMap := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	for i, j := 0, len(s)-1; i < j; {
		if _, ok := dictMap[s[i]]; ok {
			if _, ok := dictMap[s[j]]; ok {
				bs[i], bs[j] = bs[j], bs[i]
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return string(bs)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}
