package main

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	// dp[i]定义 记录 s[i-1]结尾的最长合法括号子串长度
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			// 左括号不可能是合法括号子串的结尾
			dp[i+1] = 0
		} else {
			// 遇到右括号
			if len(stack) != 0 {
				// 配对的左括号对应索引
				leftIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// 以这个右括号结尾的最长子串长度
				dp[i+1] = 1 + i - leftIndex + dp[leftIndex]
			} else {
				// 没有配对的左括号
				dp[i+1] = 0
			}
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	println(longestValidParentheses("(()"))
	println(longestValidParentheses("()(()"))
	println(longestValidParentheses(")()())"))
	println(longestValidParentheses(""))
}
