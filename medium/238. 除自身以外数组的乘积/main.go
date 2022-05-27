package main

import "fmt"

// res[i] = prefix[i-1] * suffix[n-i]  前缀乘积 * 后缀乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// 前缀乘积
	prefix := make([]int, n+1)
	// 后缀乘积
	suffix := make([]int, n+1)
	prefix[0], suffix[0] = 1, 1
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		suffix[i] = suffix[i-1] * nums[n-i]
	}

	res := make([]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = prefix[i-1] * suffix[n-i]
	}
	return res
}

func main() {
	fmt.Printf("%v\n", productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Printf("%v\n", productExceptSelf([]int{1, 2, 3, 4}))
}
