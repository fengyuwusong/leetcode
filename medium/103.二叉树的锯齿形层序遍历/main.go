package main

import "fmt"

// author: fengyuwusong date: 2022-08-04 10:24:55

// 103. 二叉树的锯齿形层序遍历
//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// Related Topics树 | 广度优先搜索 | 二叉树
//
// 👍 676, 👎 0
//
//
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	var traverse func(root *TreeNode, level int)
	traverse = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) < level+1 {
			ans = append(ans, []int{})
		}
		if level%2 == 0 {
			// 偶数层 从左到右
			ans[level] = append(ans[level], root.Val)
		} else {
			// 奇数层 从右到左
			ans[level] = append([]int{root.Val}, ans[level]...)
		}
		// 递归左右子树
		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
	}
	traverse(root, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}
