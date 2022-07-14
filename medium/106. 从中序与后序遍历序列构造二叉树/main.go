package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var f func(inorder, postorder []int) *TreeNode
	f = func(inorder, postorder []int) *TreeNode {
		if len(postorder) == 0 {
			return nil
		}

		rootVal := postorder[len(postorder)-1]
		// 遍历中序遍历得到当前根节点所在位置【此处可根据充值左右边界不循环遍历】
		i := 0
		for ; i < len(inorder); i++ {
			if inorder[i] == rootVal {
				break
			}
		}

		return &TreeNode{
			Val:   postorder[len(postorder)-1],
			Left:  f(inorder[:i], postorder[:i]),
			Right: f(inorder[i+1:], postorder[i:len(postorder)-1]),
		}
	}

	return f(inorder, postorder)
}

func main() {
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(root)
}
