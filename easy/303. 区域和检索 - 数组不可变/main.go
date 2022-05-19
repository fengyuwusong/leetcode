package main

import "fmt"

// 前缀和

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := range nums {
		preSum[i+1] = preSum[i] + nums[i]
	}
	fmt.Printf("preSum: %v\n", preSum)
	return NumArray{preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	fmt.Printf("this.preSum[right+1]: %v\n", this.preSum[right+1])
	fmt.Printf("this.preSum[left+1]: %v\n", this.preSum[left+1])
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Printf("na.SumRange(0, 2): %v\n", na.SumRange(0, 2))
}
