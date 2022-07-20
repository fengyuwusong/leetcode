package main

import "fmt"

func canJump(nums []int) bool {
	length := len(nums)
	farthest := 0
	for i := 0; i < length-1; i++ {
		// 计算在当前坐标上能跳到的最远距离
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		// 如当前坐标上能到的最远距离小于等于当前坐标，则说明没办法跳到最后一个坐标
		if farthest <= i {
			return false
		}
	}
	return farthest >= length-1
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 0, 0}))
}
