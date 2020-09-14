package main

import "fmt"

/**
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


示例：



输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。


提示：

树中的节点数小于 6000
-100 <= node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 解法1 dfs 时间复杂度O(n) 空间复杂度O(n)
// 解法2 利用已建立的next指针遍历 时间复杂度O(n) 空间复杂度O(1)
func connect(root *Node) *Node {
	n := root
	// prev 连接这层上一次成功的节点 nextLevel这层的第一个节点
	var prev, nextLevel *Node
	// 当当前节点为空且当前层处理完毕时迭代迭代
	for n != nil {
		prev, nextLevel = nil, nil
		// 循环到当前层不存在next节点
		for n != nil {
			// 连接当前节点的left节点
			connectNode(n.Left, &prev, &nextLevel)
			// 连接当前节点的right节点
			connectNode(n.Right, &prev, &nextLevel)
			// 当前节点连接完成 迭代下一个节点
			n = n.Next
		}
		// 当前层的所有子节点连接完毕 迭代下一层起始节点
		n = nextLevel
	}

	return root
}

// 此处需使用 **Node 记录存放Node指针的指针地址 由此可以对参数赋值影响原值
func connectNode(n *Node, prev, nextLevel **Node) {
	switch {
	case n == nil:
		return
	case *nextLevel == nil:
		// 记录当前层第一个不是nil的节点
		*nextLevel = n
		*prev = n
		return
	}

	// 已有当前层信息 直接用上一个已连接节点连接此节点
	(*prev).Next = n
	// 上一个已连接节点更新
	*prev = n
}

func main() {
	node1 := &Node{
		Val: 1,
	}
	node2 := &Node{
		Val: 2,
	}
	node3 := &Node{
		Val: 3,
	}
	node4 := &Node{
		Val: 4,
	}
	node5 := &Node{
		Val: 5,
	}
	node6 := &Node{
		Val: 6,
	}
	node7 := &Node{
		Val: 7,
	}
	node8 := &Node{
		Val: 8,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6
	node4.Left = node7
	node6.Right = node8
	n := connect(node1)
	fmt.Println(n)
}
