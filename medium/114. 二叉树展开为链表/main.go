package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// 后续遍历位置
	// 1. 左右子树已被拉成一条链表
	left := root.Left
	right := root.Right

	// 2. 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3. 将原先右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

func main() {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node5 := &TreeNode{
		Val: 5,
	}
	node6 := &TreeNode{
		Val: 6,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	flatten(node1)
}
