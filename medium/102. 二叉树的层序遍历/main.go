package main

import (
	"fmt"
)

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
通过次数176,208提交次数278,928

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	level := []int{0}
	ans := make([][]int, 0)
	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]
		l := level[0]
		level = level[1:]
		if len(ans) < l+1 {
			ans = append(ans, []int{root.Val})
		} else {
			ans[l] = append(ans[l], root.Val)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
			level = append(level, l+1)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
			level = append(level, l+1)
		}
	}
	return ans
}

func main() {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node1.Left = node2
	node1.Right = node3
	fmt.Println(levelOrder(node1))
}
