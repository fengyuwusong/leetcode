package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	adj []*Node
	val int
}

func amountOfTime(root *TreeNode, start int) int {
	node := findNode(rebuild(root), nil, start)
	return maxDepth(node, nil, 0)
}

func findNode(root, p *Node, val int) *Node {
	if root.val == val {
		return root
	}
	for _, v := range root.adj {
		if v == p {
			continue
		}
		ans := findNode(v, root, val)
		if ans != nil {
			return ans
		}
	}
	return nil
}

func maxDepth(root, p *Node, depth int) int {
	ans := depth
	for _, node := range root.adj {
		if node == p {
			continue
		}
		ans = max(ans, maxDepth(node, root, depth+1))
	}
	return ans
}

func addEdge(a, b *Node) {
	if a == nil || b == nil {
		return
	}
	a.adj = append(a.adj, b)
	b.adj = append(b.adj, a)
}

func rebuild(root *TreeNode) *Node {
	if root == nil {
		return nil
	}
	node := &Node{
		val: root.Val,
	}
	addEdge(node, rebuild(root.Left))
	addEdge(node, rebuild(root.Right))
	return node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 6},
		},
	}
	fmt.Println(amountOfTime(root, 3))
}
