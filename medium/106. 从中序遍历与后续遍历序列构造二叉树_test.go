package medium

import (
	"fmt"
	"testing"
)

/**
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
通过次数49,272提交次数70,958

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法1 根据hash表中序遍历节点位置递归构建数
// 解法2 迭代法 后续遍历栈的方式存储 并出栈依次对中序遍历比较得到每个节点的左右子树节点 并构建
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 递归出口
	if len(postorder) == 0 {
		return nil
	}
	// 根据hash表取得后序遍历根节点在中序遍历的位置
	rootVal := postorder[len(postorder)-1]
	// 遍历中序遍历得到当前根节点所在位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}

func TestBuildTree106(t *testing.T) {
	node := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(node)
}
