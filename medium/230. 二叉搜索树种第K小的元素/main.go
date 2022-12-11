package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历 获取第k个元素
func kthSmallest(root *TreeNode, k int) int {
	// 记录当前遍历到第几个
	var curr, res int
	var traverse func(tree *TreeNode)
	traverse = func(tree *TreeNode) {
		if tree == nil {
			return
		}
		traverse(tree.Left)
		// 中序遍历
		curr++
		if curr == k {
			res = tree.Val
			return
		}
		traverse(tree.Right)
	}
	traverse(root)
	return res
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	fmt.Println(kthSmallest(root, 1))
}
