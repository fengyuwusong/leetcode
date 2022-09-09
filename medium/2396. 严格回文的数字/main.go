package main

import (
	"fmt"
	"strconv"
)

func isStrictlyPalindromic(n int) bool {
	// 计算整数n在不同进制下的表示
	for i := 2; i <= n-2; i++ {
		s := strconv.FormatInt(int64(n), i)
		if !isPalindromic(s) {
			return false
		}
	}
	return true
}

func isPalindromic(n string) bool {
	i, j := 0, len(n)-1
	for i < j {
		if n[i] != n[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isStrictlyPalindromic(9))
	fmt.Println(isStrictlyPalindromic(4))
	fmt.Println(isStrictlyPalindromic(5))
}
