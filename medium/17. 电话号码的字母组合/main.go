package main

func letterCombinations(digits string) []string {
	mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	if len(digits) == 0 {
		return res
	}
	var backtrack func(string, int, string)
	backtrack = func(digits string, start int, s string) {
		if len(s) == len(digits) {
			// 到达回溯树底部
			res = append(res, s)
			return
		}

		// 回溯算法框架
		for i := start; i < len(digits); i++ {
			digit := digits[i] - '0'
			for _, c := range mapping[digit] {
				// 做选择
				s += string(c)
				// 递归下一层回溯树
				backtrack(digits, i+1, s)
				// 撤销选择
				s = s[:len(s)-1]
			}
		}
	}
	// 从digits[0]开始回溯
	backtrack(digits, 0, "")
	return res
}

func main() {
	for _, s := range letterCombinations("23") {
		println(s)
	}
}
