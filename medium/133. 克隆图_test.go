package medium

import (
	"fmt"
	"testing"
)

/**
给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

class Node {
    public int val;
    public List<Node> neighbors;
}


测试用例格式：

简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。

邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。

给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。
*/

//Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// 解法1：使用dfs 时间复杂度O(n) 空间复杂度O(N) 哈希表需要O(n)，栈需要O(H)
// 解法2：使用bfs 时间空间复杂度同上

func cloneGraph(node *Node) *Node {
	// hash表记录已拜访节点
	visit := make(map[*Node]*Node)
	var dfsClone func(node *Node) *Node
	dfsClone = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		// 已拜访过 直接返回已创建的节点 不在进行下面步骤
		if r, ok := visit[node]; ok {
			return r
		}

		// 未拜访 创建新节点 并根据node值对邻居节点进行dfs
		root := &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, len(node.Neighbors)),
		}

		visit[node] = root
		// 从相邻节点进行dfs
		for i, n := range node.Neighbors {
			// clone 相邻节点 并添加到成员中
			root.Neighbors[i] = dfsClone(n)
		}
		return root
	}
	return dfsClone(node)
}

func TestCloneGraph(t *testing.T) {
	node2 := &Node{
		Val:       2,
		Neighbors: nil,
	}
	node1 := &Node{
		Val:       1,
		Neighbors: nil,
	}
	node3 := &Node{
		Val:       3,
		Neighbors: nil,
	}
	node4 := &Node{
		Val:       4,
		Neighbors: nil,
	}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}
	q := cloneGraph(node1)
	fmt.Println(q)
	node := &Node{
		Val:       1,
		Neighbors: nil,
	}
	qq := cloneGraph(node)
	fmt.Println(qq)
}
