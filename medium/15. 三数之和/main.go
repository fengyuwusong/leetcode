package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	return nSumTargets(nums, 3, 0, 0)
}

func nSumTargets(nums []int, n, start, target int) [][]int {
	sz := len(nums)
	var res [][]int

	// 至少2sum 且数组大小不应该小于n
	if n < 2 || sz < n {
		return res
	}
	// 2sum base case
	if n == 2 {
		// 双指针逻辑
		lo, hi := start, sz-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				// lo小于hi 且lo对应值不重复 lo步进
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				// hi大于lo 且hi对应值不重复 hi递减
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{left, right})
				// 将相同的指针结果步进
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		// n>2 递归计算(n-1)sum结果
		for i := start; i < sz; i++ {
			sub := nSumTargets(nums, n-1, i+1, target-nums[i])
			for _, v := range sub {
				res = append(res, append(v, nums[i]))
			}
			// 去除当前重复的起始值
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

func main() {
	for _, ints := range threeSum([]int{-1, 0, 1, 2, -1, -4}) {
		for _, i := range ints {
			fmt.Printf("%d, ", i)
		}
		fmt.Println()
	}
}
