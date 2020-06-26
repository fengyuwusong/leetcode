package algorithm

import (
	"fmt"
	"testing"
)

// 算法精髓
// 计算部分匹配表 PMT
// 由PMT可得next数组的关系
// 参考文章 https://www.zhihu.com/question/21923021

func kmp(s1, s2 string) int {
	next := buildNext(s2)
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if j == -1 || s1[i] == s2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(s2) {
		return i - j
	}
	return -1
}

func buildNext(s string) []int {
	// 初始化next数组
	next := make([]int, len(s))
	// 第一位赋值为-1 由于存在 j= next[j] 如第一位赋值0则会产生死循环
	next[0] = -1
	// 主串指针
	i := 0
	j := -1 // 模式串指针 赋值-1 与上面同理
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			// 模式串处于重开状态 或 当前字符相等 当前字符前进 赋值next表
			i++
			j++
			// 重开状态 next[i]=0 即主串当前字符并不是子串开头
			next[i] = j
		} else {
			// 否则模式串起始位置回归上一个最长前缀后缀相同位置
			j = next[j]
		}
	}
	return next
}

func TestKmp(t *testing.T) {
	fmt.Println(kmp("ababababca", "abababca"))
}
