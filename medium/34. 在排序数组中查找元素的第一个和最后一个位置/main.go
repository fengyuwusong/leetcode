package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{searchLeftBound(nums, target), searchRightBound(nums, target)}
}

func searchLeftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	// target 比所有数都大
	if l == len(nums) {
		return -1
	}

	// 左边界等于 target
	if nums[l] == target {
		return l
	}
	return -1
}

func searchRightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	// 目标数比所有数都小
	if l == 0 {
		return -1
	}
	if nums[l-1] == target {
		return l - 1
	}
	return -1
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
