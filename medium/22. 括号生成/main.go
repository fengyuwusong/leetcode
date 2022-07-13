package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(track string, left, right int)
	dfs = func(track string, left, right int) {
		if len(track) == n*2 {
			res = append(res, track)
			return
		}
		if left > 0 {
			// 选择
			track += "("
			// 递归
			dfs(track, left-1, right)
			// 撤销
			track = track[:len(track)-1]
		}
		// 填充右括号只能出现在右括号比左括号多的场景
		if right > left {
			track += ")"
			dfs(track, left, right-1)
			track = track[:len(track)-1]
		}
	}
	dfs("", n, n)
	return res
}

func main() {
	for _, v := range generateParenthesis(3) {
		fmt.Println(v)
	}
}
