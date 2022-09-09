package main

import "fmt"

//给定一棵二叉树 root，返回所有重复的子树。
//
// 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
// 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,null,2,4,null,null,4]
//输出：[[2,4],[4]]
//
// 示例 2：
//
//
//
//
//输入：root = [2,1,1]
//输出：[[1]]
//
// 示例 3：
//
//
//
//
//输入：root = [2,2,2,3,null,3,null]
//输出：[[2,3],[3]]
//
//
//
// 提示：
//
//
// 树中的结点数在[1,10^4]范围内。
// -200 <= Node.val <= 200
//
// Related Topics树 | 深度优先搜索 | 哈希表 | 二叉树
//
// 👍 474, 👎 0
//
//
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// fengyuwusong 2022-09-05 01:04:45
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo := make(map[string]int)
	var backtrack func(node *TreeNode) string
	var ans []*TreeNode
	backtrack = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		lStr := backtrack(node.Left)
		rStr := backtrack(node.Right)
		val := fmt.Sprintf("l%s_m%d_r%s", lStr, node.Val, rStr)
		if nums, ok := memo[val]; ok && nums == 1 {
			ans = append(ans, node)
			memo[val]++
		} else {
			memo[val]++
		}
		return val
	}
	backtrack(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	nodes := findDuplicateSubtrees(root)
	fmt.Println(nodes)
	root = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: 0},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  0,
			Left: nil,
			Right: &TreeNode{
				Val:  0,
				Left: nil,
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: &TreeNode{Val: 0},
				},
			},
		},
	}
	nodes = findDuplicateSubtrees(root)
	fmt.Println(nodes)
}
