package main

import (
	"fmt"
	"strconv"
)

//给你一棵二叉树的根节点 root ，请你构造一个下标从 0 开始、大小为 m x n 的字符串矩阵 res ，用以表示树的 格式化布局 。构造此格式化布局矩
//阵需要遵循以下规则：
//
//
// 树的 高度 为 height ，矩阵的行数 m 应该等于 height + 1 。
// 矩阵的列数 n 应该等于 2ʰᵉⁱᵍʰᵗ⁺¹ - 1 。
// 根节点 需要放置在 顶行 的 正中间 ，对应位置为 res[0][(n-1)/2] 。
// 对于放置在矩阵中的每个节点，设对应位置为 res[r][c] ，将其左子节点放置在 res[r+1][c-2ʰᵉⁱᵍʰᵗ⁻ʳ⁻¹] ，右子节点放置在
//res[r+1][c+2ʰᵉⁱᵍʰᵗ⁻ʳ⁻¹] 。
// 继续这一过程，直到树中的所有节点都妥善放置。
// 任意空单元格都应该包含空字符串 "" 。
//
//
// 返回构造得到的矩阵 res 。
//
// 2的(4-2-1)
//
//
//
// 示例 1：
//
//
//输入：root = [1,2]
//输出：
//[["","1",""],
// ["2","",""]]
//
//
// 示例 2：
//
//
//输入：root = [1,2,3,null,4]
//输出：
//[["","","","1","","",""],
// ["","2","","","","3",""],
// ["","","4","","","",""]]
//
//
//
//
// 提示：
//
//
// 树中节点数在范围 [1, 2¹⁰] 内
// -99 <= Node.val <= 99
// 树的深度在范围 [1, 10] 内
//
// Related Topics树 | 深度优先搜索 | 广度优先搜索 | 二叉树
//
// 👍 193, 👎 0
//
//
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// fengyuwusong 2022-08-22 23:26:46
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	// 获取最大深度
	maxDepth := maxDepth(root)
	// 获取矩阵的行数
	col := 1<<maxDepth - 1
	// 初始化矩阵
	ans := make([][]string, maxDepth)
	for i := 0; i < maxDepth; i++ {
		ans[i] = make([]string, int(col))
	}

	var dfs func(root *TreeNode, row, col int)
	dfs = func(root *TreeNode, row, col int) {
		if root == nil {
			return
		}
		ans[row][col] = strconv.Itoa(root.Val)
		// dfs遍历左右子树 并计算出子树的行数 和 列数
		if root.Left != nil {
			dfs(root.Left, row+1, col-1<<(maxDepth-row-2))
		}
		if root.Right != nil {
			dfs(root.Right, row+1, col+1<<(maxDepth-row-2))
		}
	}
	dfs(root, 0, (col-1)/2)
	return ans
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
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
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(printTree(root))
}
