package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	return traverse(root)
}

func traverse(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := traverse(root.Left)
	right := traverse(root.Right)
	root.Left, root.Right = right, left
	return root
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

	println(invertTree(&t1))
}
