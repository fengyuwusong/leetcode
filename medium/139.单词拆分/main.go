package main

import (
	"fmt"
)

// author: fengyuwusong date: 2022-08-08 11:47:08

// 139. 单词拆分
//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
//
//
// 示例 1：
//
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
//
// 示例 2：
//
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//     注意，你可以重复使用字典中的单词。
//
//
// 示例 3：
//
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中的所有字符串 互不相同
//
// Related Topics字典树 | 记忆化搜索 | 哈希表 | 字符串 | 动态规划
//
// 👍 1748, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	hashTable := make(map[string]struct{})
	lenTable := make(map[int]struct{})
	// 备忘录
	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}
	for _, word := range wordDict {
		lenTable[len(word)] = struct{}{}
		hashTable[word] = struct{}{}
	}
	var check func(start int) bool
	check = func(start int) bool {
		if start == len(s) {
			return true
		}
		if memo[start] == 1 {
			return true
		} else if memo[start] == 0 {
			return false
		}
		for k, _ := range lenTable {
			if len(s)-start < k {
				continue
			}
			if _, ok := hashTable[s[start:start+k]]; !ok {
				continue
			}
			memo[start] = 1
			if check(start + k) {
				return true
			}
		}
		memo[start] = 0
		return false
	}
	return check(0)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(wordBreak("cars", []string{"car", "ca", "rs"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aa"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}
