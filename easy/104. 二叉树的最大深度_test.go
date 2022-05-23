package easy

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
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
// 后续遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1 := maxDepth(root.Left) + 1
	d2 := maxDepth(root.Right) + 1
	return max(d1, d2)
}

func TestMaxDepth(t *testing.T) {
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
	fmt.Println(maxDepth(node1))
}
