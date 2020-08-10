package medium

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
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

// 递归
func preorderTraversal(root *TreeNode) []int {
	var potFun func(root *TreeNode)
	var ans []int
	potFun = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		potFun(root.Left)
		potFun(root.Right)
	}
	potFun(root)
	return ans
}

func TestPreorderTraversal(t *testing.T) {
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
	fmt.Println(preorderTraversal(node1))
}
