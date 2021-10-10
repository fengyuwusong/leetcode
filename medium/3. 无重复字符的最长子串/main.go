package main

import (
	"fmt"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法 滑动窗口 + hash 时间复杂度O(n)
func lengthOfLongestSubstring(s string) int {
	var right, ans int
	nums := make([]int, 256)
	// 依次移动左指针 计算每个位置的
	for left := 0; left < len(s) && right < len(s); left++ {
		// 移动右指针直至出现重复值
		for right < len(s) && nums[s[right]] == 0 {
			nums[s[right]]++
			right++
		}
		// 出现重复值后计算当前长度
		if right-left > ans {
			ans = right - left
		}
		// 进入下一次遍历左指针+1 及去除左指针首字母，重新开始计算左右区间下一个无重复字符的最长子串
		nums[s[left]]--
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
