package main

import (
	"fmt"
	"math"
)

//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不
//一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//
//
// 提示：
//
//
// 树中节点数目范围是 [1, 3 * 10⁴]
// -1000 <= Node.val <= 1000
//
// Related Topics树 | 深度优先搜索 | 动态规划 | 二叉树
//
// 👍 1675, 👎 0
//
//
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2022-08-09 20:15:18 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	// 后续遍历 获取当前节点的最大单边路径和
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, traverse(root.Left))
		right := max(0, traverse(root.Right))
		// 当前节点的最大路径和
		sum := left + right + root.Val
		if sum > ans {
			ans = sum
		}
		return max(left, right) + root.Val
	}
	traverse(root)
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
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(maxPathSum(root))
	root = &TreeNode{
		Val:   -10,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(maxPathSum(root))
	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(maxPathSum(root))
}
