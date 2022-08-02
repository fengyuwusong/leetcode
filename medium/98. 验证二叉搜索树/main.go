package main

import (
	"fmt"
)

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
//
//
// 示例 1：
//
//
//输入：root = [2,1,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
//
//
//
// 提示：
//
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
// Related Topics树 | 深度优先搜索 | 二叉搜索树 | 二叉树
//
// 👍 1695, 👎 0
//
//
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2022-08-03 01:16:16 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	var helper func(root, min, max *TreeNode) bool
	helper = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		// 递归检查左右子树
		return helper(root.Left, min, root) && helper(root.Right, root, max)
	}

	return helper(root, nil, nil)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}
	fmt.Println(isValidBST(root))
	root = &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(isValidBST(root))
}
