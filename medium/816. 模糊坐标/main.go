package main

import (
	"fmt"
)

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	var ans []string
	for i := 0; i < len(s)-1; i++ { // 枚举逗号 在i的后面追加逗号
		a := search(s, 0, i)
		b := search(s, i+1, len(s)-1)
		for _, x := range a {
			for _, y := range b {
				ans = append(ans, fmt.Sprintf("(%s, %s)", x, y))
			}
		}
	}
	return ans
}

// 该字符可以组成的数值数组组合
func search(s string, start, end int) []string {
	var ans []string
	if start == end || s[start] != '0' {
		ans = append(ans, s[start:end+1])
	}
	for i := start; i < end; i++ { // 枚举小数点 在i后面追加小数点
		a := s[start : i+1]
		b := s[i+1 : end+1]
		if len(a) > 1 && a[0] == '0' {
			continue
		}
		if b[len(b)-1] == '0' {
			continue
		}
		ans = append(ans, a+"."+b)
	}
	return ans
}

func main() {
	fmt.Println(ambiguousCoordinates("(123)"))
}
