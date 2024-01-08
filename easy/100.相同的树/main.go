package main

import "fmt"

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
// 示例 1：
//
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
//
// 示例 2：
//
// 输入：p = [1,2], q = [1,null,2]
// 输出：false
//
// 示例 3：
//
// 输入：p = [1,2,1], q = [1,1,2]
// 输出：false
//
// 提示：
//
// 两棵树上的节点数目都在范围 [0, 100] 内
// -10⁴ <= Node.val <= 10⁴
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1114 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// author: fengyuwusong
// create time: 2024-01-09 00:25:43
// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, p, q)
	for len(queue) > 0 {
		p = queue[0]
		q = queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left, q.Left)
		queue = append(queue, p.Right, q.Right)
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	p := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	p1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: nil,
	}
	q1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{Val: 2},
	}

	p2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	q2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(p, q))
	fmt.Println(isSameTree(p1, q1))
	fmt.Println(isSameTree(p2, q2))
}
