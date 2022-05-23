package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	l []string // 记录dfs后每个节点信息
}

func Constructor() Codec {
	return Codec{}
}

// 解法1 dfs/bfs 序列化后反序列化
// 树 1,2,3,nil,nil,4,5 dfs序列化后=> 1,2,nil,nil,3,4,nil,nil,5,nil,nil
// 反序列化则为先解析左字数，在解析右子树

// 解法2 使用括号标识编码 + 递归下降编码
// 定义二叉树 => (<left_tree>)cur_num(<right_tree>)  空为X
// 那么 根据该定义序列化树 1,2,3,nil,nil,4,5 => ((X)2(X))1((4)3(5))
// 那么由定义可以推导出反序列化式BNF  T->(T) num (T) | X
// 用 T 代表一棵树序列化之后的结果，| 表示 T 的构成为 (T) num (T) 或者 X，| 左边是对 T 的递归定义，右边规定了递归终止的边界条件。
// 我们只需要通过开头的第一个字母是 X 还是 ( 来判断使用哪一种解析方法
// 例如上述 ((X)2(X))1((4)3(5)) 由递归符合 (T)num(T)=> ((X)2(X)) 1 ((4)3(5)) 再分别对左右子书递归即可得到树结构

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// dfs
	if root == nil {
		return "nil"
	}
	return fmt.Sprintf("%s,%s,%s", strconv.Itoa(root.Val), this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodeList := strings.Split(data, ",")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		// 出栈
		node := nodeList[0]
		nodeList = nodeList[1:]
		if node == "nil" {
			return nil
		}
		// 构建根节点
		val, _ := strconv.Atoi(node)
		root := &TreeNode{Val: val}
		// 递归生成左子树
		root.Left = dfs()
		// 递归生成右子树
		root.Right = dfs()
		return root
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

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
	node3.Left = node4
	node3.Right = node5
	node4.Left = node6
	node4.Right = node7
	c := Constructor()
	ans := c.serialize(node1)
	fmt.Println(ans)
	t := c.deserialize(ans)
	fmt.Println(t)

	ans2 := c.serialize(nil)
	t2 := c.deserialize(ans2)
	fmt.Println(t2)
}
