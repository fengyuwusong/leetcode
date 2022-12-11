package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var traverse func(tree *TreeNode)
	var temp int
	traverse = func(tree *TreeNode) {
		if tree == nil {
			return
		}
		traverse(tree.Right)
		if temp != 0 {
			tree.Val += temp
		}
		temp = tree.Val
		traverse(tree.Left)
	}
	traverse(root)
	return root
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: &TreeNode{Val: 3},
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: &TreeNode{Val: 5},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: &TreeNode{Val: 8},
			},
		},
	}
	convertBST(root)
	fmt.Println(root)
}
