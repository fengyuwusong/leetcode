package main

import "fmt"

//给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
//
// 树的 最大宽度 是所有层中最大的 宽度 。
//
//
//
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的
//null 节点，这些 null 节点也计入长度。
//
// 题目数据保证答案将会在 32 位 带符号整数范围内。
//
//
//
// 示例 1：
//
//
//输入：root = [1,3,2,5,3,null,9]
//输出：4
//解释：最大宽度出现在树的第 3 层，宽度为 4 (5,3,null,9) 。
//
//
// 示例 2：
//
//
//输入：root = [1,3,2,5,null,null,9,6,null,7]
//输出：7
//解释：最大宽度出现在树的第 4 层，宽度为 7 (6,null,null,null,null,null,7) 。
//
//
// 示例 3：
//
//
//输入：root = [1,3,2,5]
//输出：2
//解释：最大宽度出现在树的第 2 层，宽度为 2 (3,2) 。
//
//
//
//
// 提示：
//
//
// 树中节点的数目范围是 [1, 3000]
// -100 <= Node.val <= 100
//
//
//
// Related Topics树 | 深度优先搜索 | 广度优先搜索 | 二叉树
//
// 👍 406, 👎 0
//
//
//
//

// fengyuwusong 2022-08-23 00:38:48
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
func widthOfBinaryTree(root *TreeNode) int {
	var ans int
	var levelMemo [][]int
	var dfs func(tree *TreeNode, level, index int)
	dfs = func(tree *TreeNode, level, index int) {
		if tree == nil {
			return
		}
		if len(levelMemo) < level {
			levelMemo = append(levelMemo, []int{index, index})
		}
		// 扩大当前节点的最大最小值 并计算答案
		if index < levelMemo[level-1][0] {
			levelMemo[level-1][0] = index
		}
		if index > levelMemo[level-1][1] {
			levelMemo[level-1][1] = index
		}
		if levelMemo[level-1][1]-levelMemo[level-1][0]+1 > ans {
			ans = levelMemo[level-1][1] - levelMemo[level-1][0] + 1
		}
		dfs(tree.Left, level+1, index<<1)
		dfs(tree.Right, level+1, index<<1+1)
	}
	dfs(root, 1, 1)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 6},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  9,
				Left: &TreeNode{Val: 7},
			},
		},
	}
	fmt.Println(widthOfBinaryTree(root))
}
