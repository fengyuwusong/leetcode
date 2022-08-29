package main

import "fmt"

func largestPalindromic(num string) string {
	arr := make([]int, 10)
	for i := 0; i < len(num); i++ {
		arr[num[i]-'0']++
	}
	var ans string
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			n := arr[i] / 2
			for j := 0; j < n; j++ {
				ans = fmt.Sprintf("%d%s%d", i, ans, i)
			}
		}
	}
	for i := 9; i >= 0; i-- {
		if arr[i] > 0 && arr[i]%2 == 1 {
			ans = fmt.Sprintf("%s%d%s", ans[:len(ans)/2], i, ans[len(ans)/2:])
			break
		}
	}
	// 去除前置0
	for i := 0; i < len(ans); i++ {
		if ans[i] != '0' {
			return ans[i : len(ans)-i]
		}
	}
	if ans[0] == '0' {
		return "0"
	}
	return ans
}

func main() {
	fmt.Println(largestPalindromic("444947137"))
	fmt.Println(largestPalindromic("00009"))
	fmt.Println(largestPalindromic("00001105")) //1005001
	fmt.Println(largestPalindromic("00000"))    //0
}
