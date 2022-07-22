package main

import "fmt"

// author: fengyuwusong date: 2022-07-22 15:57:34

// 72. 编辑距离
//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
// Related Topics字符串 | 动态规划
//
// 👍 2493, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	// 初始化dp数组 dp[i][j] 为 word1[0..i] 和 word2[0..j] 的最小编辑距离
	// dp[0][0..j] = word1[0..j] 的最小编辑距离 即 0-j(插入j个字符一致 j次编辑记录)
	// dp[i][0] 同理
	// 其中dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
	// 即为上一次可能字符的到达dp[i][j]的最短距离(word1, word2分别各差一个字符和一起差一个字符) + 1 (删除一个字符，插入一个字符，替换一个字符)
	// 如果word1[i] == word2[j] 则dp[i][j] = dp[i-1][j-1]
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		}
		return k
	}
	if j < k {
		return j
	}
	return k
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
