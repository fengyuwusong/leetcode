package main

import "fmt"

func findSubarrays(nums []int) bool {
	memo := make(map[int]bool)
	for i := 1; i < len(nums); i++ {
		sum := nums[i] + nums[i-1]
		if memo[sum] {
			return true
		}
		memo[sum] = true
	}
	return false
}

func main() {
	fmt.Println(findSubarrays([]int{4, 2, 4}))
	fmt.Println(findSubarrays([]int{1, 2, 3, 4, 5}))
	fmt.Println(findSubarrays([]int{0, 0, 0}))
	fmt.Println(findSubarrays([]int{1, 2, 3, 1, 0, 2}))
	fmt.Println(findSubarrays([]int{1, 1}))
}
