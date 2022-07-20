package main

import (
	"fmt"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心法 实际也是dp
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = maxInt(nums[i], nums[i]+sum)
		max = maxInt(sum, max)
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1}))
}
