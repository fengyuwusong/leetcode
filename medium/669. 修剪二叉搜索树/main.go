package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		lNode := dfs(node.Left)
		rNode := dfs(node.Right)
		// 当前节点小于阈值 则返回大于当前节点的右节点
		if node.Val < low {
			return rNode
		}
		// 反之返回左节点
		if node.Val > high {
			return lNode
		}
		// 当前节点符合条件
		node.Left = lNode
		node.Right = rNode
		return node
	}
	return dfs(root)
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  0,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: nil,
			},
		},
		Right: &TreeNode{Val: 4},
	}
	root = trimBST(root, 1, 3)
	fmt.Println(root)
}
