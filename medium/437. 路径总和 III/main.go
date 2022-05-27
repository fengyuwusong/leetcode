package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var _pathSum, _targetSum, res int
var preSumCount map[int]int

// 构建从根节点下来所有数组
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res = 0
	preSumCount = make(map[int]int)
	preSumCount[0] = 1
	_targetSum = targetSum
	// 遍历树构建前缀和数组
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	// 前序遍历位置
	_pathSum += root.Val
	// 从二叉树根节点开始 路径和 = _pathSum - targetSum 的条数
	// 就是路径和为 targetSum 的路径条数
	res += preSumCount[_pathSum-_targetSum]
	// 记录从二叉树的根节点开始 路径和  = _pathSum 的路径条数
	preSumCount[_pathSum]++

	traverse(root.Left)
	traverse(root.Right)

	// 后续遍历位置
	preSumCount[_pathSum]--
	_pathSum -= root.Val
}

func main() {
	node3 := TreeNode{Val: 3}
	node_2 := TreeNode{Val: -2}
	node3_ := TreeNode{Val: 3, Left: &node3, Right: &node_2}
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2, Right: &node1}
	node5 := TreeNode{Val: 5, Left: &node3_, Right: &node2}
	node11 := TreeNode{Val: 11}
	node_3 := TreeNode{Val: -3, Right: &node11}
	node10 := TreeNode{Val: 10, Left: &node5, Right: &node_3}

	println(pathSum(&node10, 8))

	r := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}

	println(pathSum(&r, 22))
}
