package hard

import (
	"math"
	"testing"
)

/**
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func minWindow(s string, t string) string {
	need, window := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	var left, right, valid int
	start, length := 0, math.MaxInt

	for right < len(s) {
		r := s[right]
		// 扩大窗口
		right++

		// 记录扩大窗口后 window变化
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}
		// 符合包含子串条件 收缩窗口
		for valid == len(need) {
			// 更新最小覆盖子串长度
			if right-left < length {
				start = left
				length = right - left
			}
			// 收缩窗口
			d := s[left]
			left++
			// d 存在子串中 且 窗口 d个数等于need个数 则 valid-1
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

func TestMinWindow(t *testing.T) {
	println(minWindow("ADOBECODEBANC", "ABC"))
}
