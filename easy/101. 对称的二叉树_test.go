package easy

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false


限制：

0 <= 节点个数 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	if leftNode.Val != rightNode.Val {
		return false
	}
	return checkSymmetric(leftNode.Left, rightNode.Right) && checkSymmetric(leftNode.Right, rightNode.Left)
}

func TestIsSymmetric(t *testing.T) {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 2,
	}
	node4 := &TreeNode{
		Val: 3,
	}
	node5 := &TreeNode{
		Val: 4,
	}
	node6 := &TreeNode{
		Val: 4,
	}
	node7 := &TreeNode{
		Val: 3,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	fmt.Println(isSymmetric(node1))
}
