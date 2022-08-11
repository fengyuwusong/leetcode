package main

import "fmt"

// author: fengyuwusong date: 2022-08-04 17:56:59

// 131. 分割回文串
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
// Related Topics字符串 | 动态规划 | 回溯
//
// 👍 1221, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
// 获取所有组合 逐个判断是否回文串
func partition(s string) [][]string {
	var ans [][]string
	var backtrack func(start int, cur []string)
	memo := make(map[string]int)
	backtrack = func(start int, cur []string) {
		if start == len(s) {
			temp := make([]string, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		for j := start; j < len(s); j++ {
			// 优先查找备忘录
			if v, ok := memo[s[start:j+1]]; ok {
				if v == 1 {
					cur = append(cur, s[start:j+1])
					backtrack(j+1, cur)
					cur = cur[:len(cur)-1]
				}
				continue
			}
			memo[s[start:j+1]] = 0
			if isPalindrome(s[start : j+1]) {
				memo[s[start:j+1]] = 1
				cur = append(cur, s[start:j+1])
				backtrack(j+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0, []string{})
	return ans
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
