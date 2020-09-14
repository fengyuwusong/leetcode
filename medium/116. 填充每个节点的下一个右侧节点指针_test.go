package medium

import (
	"fmt"
	"testing"
)

/**
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



示例：



输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}

输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}

解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。


提示：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

// 利用上层next节点连接进行递归
func connect(root *Node116) *Node116 {
	var dfs func(node *Node116, next *Node116)
	dfs = func(node *Node116, next *Node116) {
		if node != nil {
			node.Next = next
			dfs(node.Left, node.Right)
			// 通过上级next连接相邻父节点之间的子节点
			if node.Next != nil {
				dfs(node.Right, node.Next.Left)
			} else {
				// 即使next为空也仍需遍历 因为右节点下的左右节点也需要连接
				dfs(node.Right, nil)
			}
		}
	}
	dfs(root, nil)
	return root
}

func TestConnect(t *testing.T) {
	node1 := &Node116{
		Val: 1,
	}
	node2 := &Node116{
		Val: 2,
	}
	node3 := &Node116{
		Val: 3,
	}
	node4 := &Node116{
		Val: 4,
	}
	node5 := &Node116{
		Val: 5,
	}
	node6 := &Node116{
		Val: 6,
	}
	node7 := &Node116{
		Val: 7,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	n := connect(node1)
	fmt.Println(n)
}
