package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	// 计算前缀和
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	ans := make([]int, len(queries))
	j := 0
	for i := range queries {
		for j < len(nums) && nums[j] <= queries[i] {
			j++
		}
		ans[i] = j
		j = 0
	}
	return ans
}

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
	fmt.Println(answerQueries([]int{2, 3, 4, 5}, []int{1}))
	fmt.Println(answerQueries([]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}))
}
