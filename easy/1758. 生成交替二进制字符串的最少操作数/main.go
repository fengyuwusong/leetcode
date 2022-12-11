package main

import "fmt"

// 构造法 只有两种情况 从0开始和从1开始 计算出一种即可推出另一种
// 0110 => 0101 需两步 1010 需两步 4-2(前面的两步)
func minOperations(s string) int {
	cnt := 0
	for i, c := range s {
		if i%2 != int(c-'0') {
			cnt++
		}
	}
	return min(cnt, len(s)-cnt)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minOperations("110"))
	fmt.Println(minOperations("111"))
}
