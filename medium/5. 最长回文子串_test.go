package medium

import (
	"fmt"
	"testing"
)

// 暴力解 以每个字符为起点 遍历每种子串的结果并在对子串进行遍历判断是否回文 时间复杂度 O(n^3) 空间复杂度 O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	}
	iRes, jRes := 0, 1
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if isPalindrome(s, i, j) {
				if j+1-i > jRes-iRes {
					iRes, jRes = i, j+1
				}
			}
		}
	}
	return s[iRes:jRes]
}

func isPalindrome(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		return false
	}
	return true
}

// todo 中间扩散法 时间复杂度 O(n * 2n) => O(n^2) （奇偶两种情况） 空间复杂度 O（1） 由中间向两头扩散求解
//  仅需从以每个字符为中心，分别向两边扩散一次取最大值即可

// todo 动态规划 时间复杂度O(n^2) 空间复杂度 O(n^2)
//  定义子串状态 dp[i][j] 为子串 s[i:j+1]
//  那么 可以得到状态转移方程： dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
//  (即子串是否回文 = 两端字符是否相等 且 去掉两端后的字符是否回文决定)
//  边界值: j-1 - (i+1) < 2 => j-i <3时，状态转移方程不成立 （即当s[i:j+1] 长度为2或者3时不必检查是否回文）
//  当  j-i <3 时 是否回文由 s[i] = s[j] 决定
//  动态规划实际为填写二维表格

// todo manacher 算法 时间复杂度O(n) 空间复杂度O(1)
//  对原始字符串进行预处理

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("eabcb"))
	fmt.Println(longestPalindrome("abacab"))
	fmt.Println(longestPalindrome("jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"))
}
