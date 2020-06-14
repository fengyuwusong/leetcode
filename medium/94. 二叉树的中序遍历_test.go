package medium

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// dfs递归 时间复杂度O(n) 空间复杂度 最坏情况O(n)全是左节点 平均情况 O(log n)
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

// dfs 栈实现
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	cur := root
	// 当前节点不为空或栈不为空继续进行dfs
	for cur != nil || len(stack) != 0 {
		// 将当前节点的所有左节点依次入栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 取栈中最后一个值并出栈
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 将最后的左节点依次保存
		res = append(res, cur.Val)
		// 当前节点左节点遍历完毕 取当前节点的右节点
		cur = cur.Right
	}
	return res
}

func TestInorderTraversal(t *testing.T) {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node1.Right = node2
	node2.Left = node3
	fmt.Println(inorderTraversal(node1))
	fmt.Println(inorderTraversal(node3))
	fmt.Println(inorderTraversal2(node1))
	fmt.Println(inorderTraversal2(node3))
}
