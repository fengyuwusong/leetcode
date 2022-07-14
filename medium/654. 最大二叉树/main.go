package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums)
}

// 查找数组最大值
func construct(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 获取最大值数组索引
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}

	left := construct(nums[:max])
	var right *TreeNode
	if max+1 < len(nums) {
		right = construct(nums[max+1:])
	}

	root := &TreeNode{
		Val:   nums[max],
		Left:  left,
		Right: right,
	}
	return root
}

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(root)
}
