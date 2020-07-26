package easy

import (
	"fmt"
	"testing"
)

/**
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:

输入: s = "egg", t = "add"
输出: true
示例 2:

输入: s = "foo", t = "bar"
输出: false
示例 3:

输入: s = "paper", t = "title"
输出: true
说明:
你可以假设 s 和 t 具有相同的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/isomorphic-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isIsomorphic(s string, t string) bool {
	// 需使用两个hashMap分别记录两个串的映射关系
	hashMapS := make(map[byte]byte)
	hashMapT := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		v1, ok1 := hashMapS[s[i]]
		_, ok2 := hashMapT[t[i]]
		if ok1 && ok2 {
			if t[i] != v1 {
				return false
			}
		} else if !ok1 && !ok2 {
			hashMapS[s[i]], hashMapT[t[i]] = t[i], s[i]
		} else {
			return false
		}
	}
	return true
}

func TestIsIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	// 特殊情况用例 同一个字符不能被不同字符同时映射
	fmt.Println(isIsomorphic("ab", "aa"))
	fmt.Println(isIsomorphic("ab", "ca"))
}
