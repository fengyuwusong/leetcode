package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。

编码的字符串应尽可能紧凑。

示例 1：

输入：root = [2,1,3]
输出：[2,1,3]
示例 2：

输入：root = []
输出：[]

提示：

树中节点数范围是 [0, 104]
0 <= Node.val <= 104
题目数据 保证 输入的树是一棵二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/serialize-and-deserialize-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 由于二叉搜索数树的中序遍历是有序的，故可在序列化时使用前序序列，反序列化则根据前序序列和中序序列可反推出二叉树

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 构造前序遍历
	var potFun func(root *TreeNode)
	var res string
	potFun = func(root *TreeNode) {
		// 递归出口
		if root == nil {
			return
		}
		res += "," + strconv.Itoa(root.Val)
		potFun(root.Left)
		potFun(root.Right)
	}
	potFun(root)

	if len(res) == 0 {
		return ""
	}
	return res[1:]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	// 根据前序遍历获取中序遍历
	ss := strings.Split(data, ",")
	var front, middle []int
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		front = append(front, i)
		middle = append(middle, i)
	}
	sort.Ints(middle)
	return buildTree(front, middle)
}

// 利用hash表后时间复杂度O(n) 空间复杂度O(n) = 答案所需空间n + hash表空间 n + 递归所需栈临时空间 h 树高 (h<n)
// todo 解法2: 迭代法 可使用栈对前序遍历进行入栈
//  并依次对栈顶和 index 节点(当前根节点 由中序遍历决定) 进行比较
//  不相等相等则为当前根节点的左儿子
//  相等则为当前节点的右儿子 并弹出重取当前根节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	hashData := make(map[int]int, len(preorder))
	// 构建中序序列hash表
	for i, v := range inorder {
		hashData[v] = i
	}

	// 定义递归方法由前序序列拆分树左右节点递归分别取得各个节点
	bTree := func(preorder []int, inorder []int) *TreeNode {
		// 递归出口
		if len(preorder) == 0 {
			return nil
		}
		// 根据hash表取得根节点在中序遍历的位置
		rootIndex := hashData[preorder[0]]
		// 根据根节点在中序遍历的位置将前序遍历和中序遍历拆分为左右节点两半 进行递归
		// preorder[1:rootIndex+1] => 去除根节点到中序遍历位置(长度) 其他区域类似 可画图观察
		return &TreeNode{
			Val:   preorder[0],
			Left:  buildTree(preorder[1:rootIndex+1], inorder[:rootIndex]),
			Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
		}
	}
	return bTree(preorder, inorder)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

func main() {
	ser := Constructor()
	deser := Constructor()
	leftNode := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	rightNode := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   4,
		Left:  &leftNode,
		Right: &rightNode,
	}
	tree := ser.serialize(&root)
	fmt.Println(tree)
	ans := deser.deserialize(tree)
	fmt.Println(ans)

	ans2 := deser.deserialize("")
	fmt.Println(ans2)
}
