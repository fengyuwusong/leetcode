package main

import (
	"fmt"
)

func minElements(nums []int, limit int, goal int) int {
	// 计算当前总和
	var sum int
	for _, num := range nums {
		sum += num
	}
	ans := abs(goal-sum) / limit
	if abs(goal-sum)%limit != 0 {
		return ans + 1
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4))
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0))
}
