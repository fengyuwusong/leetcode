package main

import "fmt"

/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]





示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4

输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路 可先将可能出现的情况列举
// 1. 当节点node等于p || q时 node 即p q 的最近公共祖先节点
// 2. 当节点node不等于 p||q 时，有以下三种情况
// 2.1 节点node的左右子树节点分别包含 p 和 q，那么node 即p q 的最近公共祖先节点
// 2.2 节点node的左子树包含p和q，那么 node的左子树即q p的最近公共祖先节点
// 2.3 右子树同理
// 2.4 左右子树仅包含其中一个节点 q p，说明公共祖先在父节点上

// 例如树 [1, 2, 3, 4, 5, 6, 7] p = 4 q = 7
// 递归次数:
// 1. root = 1, root.val!=q p  递归左右子树
// 2.1 root = 2 root.val!=q p  递归左右子树
// 2.2 root = 3 root.val!=q p 递归左右子树
// 3.1 root = 4 root.val=p 返回4到2.1递归
// 3.2 root = 5 root.val!=q p 递归左右子树 左右子树为nil 返回nil
// 3.3/3.4 同理 返回 7到2.2
// 2.1.1 符合 2.4 返回 4 到1
// 2.2.1 符合 2.4 返回 7 到1
// 1.1 符合2.1 root = 1 为最近公共节点

// 解法1 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}

	return right
}

// 解法2 存储父节点
// 1. 使用dfs遍历整颗二叉树 用hash表记录每个节点的父节点
// 2. 从p节点开始不断向祖先移动 记录已访问过的祖先节点
// 3. q节点同理，当遇到已访问过的祖先节点时，则意味该节点时q和p的深度最深公共祖先

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
	fmt.Println(lowestCommonAncestor(node1, node4, node7).Val)
}
