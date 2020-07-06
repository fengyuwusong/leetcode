package easy

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先遍历
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Right, sum-root.Val) || hasPathSum(root.Left, sum-root.Val)
}

func TestHasPathSum(t *testing.T) {
	t7 := &TreeNode{Val: 7}
	t8 := &TreeNode{Val: 2}
	t9 := &TreeNode{Val: 1}
	t4 := &TreeNode{Val: 11, Left: t7, Right: t8}
	t5 := &TreeNode{Val: 13}
	t6 := &TreeNode{Val: 1, Right: t9}
	t2 := &TreeNode{Val: 4, Left: t4}
	t3 := &TreeNode{Val: 8, Left: t5, Right: t6}
	t1 := &TreeNode{Val: 5, Left: t2, Right: t3}

	fmt.Println(hasPathSum(t1, 22))
	t4.Val = 1
	fmt.Println(hasPathSum(t1, 22))
}
