package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	traverse(root)
	return maxDiameter
}

func traverse(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	leftMax := traverse(tree.Left)
	rightMax := traverse(tree.Right)

	// 后序位置 计算该节点直径 和 该节点的深度
	md := leftMax + rightMax
	if md > maxDiameter {
		maxDiameter = md
	}
	if leftMax > rightMax {
		return leftMax + 1
	}
	return 1 + rightMax
}

func main() {
	t5 := TreeNode{Val: 5}
	t4 := TreeNode{Val: 4}
	t2 := TreeNode{
		Val:   2,
		Left:  &t4,
		Right: &t5,
	}
	t3 := TreeNode{Val: 3}
	t1 := TreeNode{
		Val:   1,
		Left:  &t2,
		Right: &t3,
	}

	println(diameterOfBinaryTree(&t1))
}
