package medium

import (
	"fmt"
	"testing"
)

// 暴力解 以每个字符为起点 遍历每种子串的结果并在对子串进行遍历判断是否回文
// 时间复杂度 O(n^3) 空间复杂度 O(1)
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

// 解法2 中间扩散法 时间复杂度 O(n * 2n) => O(n^2) （奇偶两种情况） 空间复杂度 O（1） 由中间向两头扩散求解
//  仅需从以每个字符为中心，分别向两边扩散一次取最大值即可
func longestPalindrome2(s string) string {
	var ans string
	for index := range s {
		// 回文子串奇数情况 排除本身两边扩散
		for i, j := index-1, index+1; ; i, j = i-1, j+1 {
			if i < 0 || j >= len(s) || s[i] != s[j] {
				if j-i-1 > len(ans) {
					ans = s[i+1 : j]
				}
				break
			}
		}
		// 回文子串偶数情况 不排除本身两边扩散
		for i, j := index, index+1; ; i, j = i-1, j+1 {
			if i < 0 || j >= len(s) || s[i] != s[j] {
				if j-i-1 > len(ans) {
					ans = s[i+1 : j]
				}
				break
			}
		}
	}
	return ans
}

// 解法3 动态规划 时间复杂度O(n^2) 空间复杂度 O(n^2)
//  定义子串状态 dp[i][j] 为子串 s[i:j+1]
//  那么 可以得到状态转移方程： dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
//  (即子串是否回文 = 两端字符是否相等 且 去掉两端后的字符是否回文决定)
//  边界值: j-1 - (i+1) < 2 => j-i <3时，状态转移方程不成立 （即当s[i:j+1] 长度小于3时不必检查是否回文）
//  当  j-i =2 时即子串长度=2 是否回文由 s[i] = s[j] 决定
//  当j-i=1时即子串长度=1时 dp[i][j]=1 即本身即是回文
//  动态规划实际为填写二维表格
func longestPalindrome3(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return s
	}
	var ans string
	// 初始化dp dp[i][j]=1表示子串s[i:j]是回文子串
	dp := make([][]int, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]int, sLen)
	}
	// 开始填表
	// 按照子串长度从小到大进行遍历 子串长度长的依赖长度短的进行填表
	for l := 0; l < sLen; l++ {
		// 长度为l 起始位置为i的所有子串遍历
		for i := 0; i+l < sLen; i++ {
			j := i + l
			// 填表
			if l == 0 {
				// 当仅一个字符时 回文串长度=1
				dp[i][j] = 1
			} else if l == 1 {
				// 当两个字符时 dp[i][j] = s[i] == s[j]
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				// 当回文子串大于三个字符时 dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 当前回文子串长度填表完毕 计算当前结果值
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

// todo manacher 算法 时间复杂度O(n) 空间复杂度O(1)
//  对原始字符串进行预处理

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("eabcb"))
	fmt.Println(longestPalindrome("abacab"))
	fmt.Println(longestPalindrome("jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"))
	fmt.Println(longestPalindrome2("babad"))
	fmt.Println(longestPalindrome2("cbbd"))
	fmt.Println(longestPalindrome2("eabcb"))
	fmt.Println(longestPalindrome2("abacab"))
	fmt.Println(longestPalindrome2("jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"))
	fmt.Println(longestPalindrome3("babad"))
	fmt.Println(longestPalindrome3("cbbd"))
	fmt.Println(longestPalindrome3("eabcb"))
	fmt.Println(longestPalindrome3("abacab"))
	fmt.Println(longestPalindrome3("jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"))
}
