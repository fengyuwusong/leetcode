package medium

import (
	"fmt"
	"testing"
)

/**
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func TestBuildTree(t *testing.T) {
	//fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	node := buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{3, 2, 5, 4, 1, 7, 6, 8})
	fmt.Println(node)
}
