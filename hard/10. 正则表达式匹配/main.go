package main

// 备忘录
var memo [][]int

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	memo = make([][]int, ls)
	for i := 0; i < ls; i++ {
		memo[i] = make([]int, lp)
		for j := 0; j < lp; j++ {
			memo[i][j] = -1
		}
	}
	return dp(s, 0, p, 0)
}

func dp(s string, i int, p string, j int) bool {
	m, n := len(s), len(p)

	// base case
	if j == n {
		return i == m
	}
	if i == m {
		// 如果p的最后字符是[x*]组合，那么可以匹配0个字符
		if (n-j)%2 == 1 {
			return false
		}
		for ; j+1 < n; j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}

	// 使用备忘录
	if memo[i][j] != -1 {
		return memo[i][j] == 1
	}
	var res bool
	if s[i] == p[j] || p[j] == '.' {
		// 模式串并非最后一个字符 && 下一个字符是*
		if j < n-1 && p[j+1] == '*' {
			// 忽略下一个*字符 || 匹配串前进
			res = dp(s, i, p, j+2) || dp(s, i+1, p, j)
		} else {
			// 两者同时步进
			res = dp(s, i+1, p, j+1)
		}
	} else {
		if j < n-1 && p[j+1] == '*' {
			// 忽略当前字符及下一个* 模式串前进
			res = dp(s, i, p, j+2)
		} else {
			res = false
		}
	}

	switch res {
	case true:
		memo[i][j] = 1
	case false:
		memo[i][j] = 0
	}
	return res
}

func main() {
	println(isMatch("aa", "a"))
	println(isMatch("aa", "a*"))
	println(isMatch("ab", ".*"))
}
