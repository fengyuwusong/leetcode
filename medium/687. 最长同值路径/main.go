package main

import "fmt"

//给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
//
// 两个节点之间的路径长度 由它们之间的边数表示。
//
//
//
// 示例 1:
//
//
//
//
//输入：root = [5,4,5,1,1,5]
//输出：2
//
//
// 示例 2:
//
//
//
//
//输入：root = [1,4,5,4,4,5]
//输出：2
//
//
//
//
// 提示:
//
//
// 树的节点数的范围是 [0, 10⁴]
// -1000 <= Node.val <= 1000
// 树的深度将不超过 1000
//
// Related Topics树 | 深度优先搜索 | 二叉树
//
// 👍 595, 👎 0
//
//
//
//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// fengyuwusong 2022-09-02 00:42:50
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lNums := dfs(node.Left)
		rNums := dfs(node.Right)
		// 当前节点的左右分支边数
		l1, r1 := 0, 0
		// 如左节点和当前节点值相同 则左节点的边数等于子节点的边数+1
		if node.Left != nil && node.Left.Val == node.Val {
			l1 = lNums + 1
		}
		// 同上
		if node.Right != nil && node.Right.Val == node.Val {
			r1 = rNums + 1
		}
		// 当前节点的边数等于左右节点边数的和
		ans = max(ans, l1+r1)
		// 返回当前节点的单边的边数
		return max(l1, r1)
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(longestUnivaluePath(root))
}
