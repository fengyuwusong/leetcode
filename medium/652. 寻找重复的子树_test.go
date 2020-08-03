package medium

import (
	"fmt"
	"strconv"
	"testing"
)

/**
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

两棵树重复是指它们具有相同的结构以及相同的结点值。

示例 1：

        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4
下面是两个重复的子树：

      2
     /
    4
和

    4
因此，你需要以列表的形式返回上述重复子树的根结点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-duplicate-subtrees
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

// 使用dfs的方式遍历数得到所有子树数据并记录hash 重复则加入结果数组返回
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	hashMap := make(map[string]int)
	var ans []*TreeNode
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		res := strconv.Itoa(root.Val)
		res += "l" + dfs(root.Left)
		res += "r" + dfs(root.Right)
		if d, ok := hashMap[res]; ok && d == 1 {
			ans = append(ans, root)
		}
		hashMap[res]++
		return res
	}
	dfs(root)
	return ans
}

func TestFindDuplicateSubtrees(t *testing.T) {
	t7 := &TreeNode{
		Val: 4,
	}
	t6 := &TreeNode{Val: 4}
	t5 := &TreeNode{
		Val:  2,
		Left: t7,
	}
	t4 := &TreeNode{Val: 4}
	t3 := &TreeNode{
		Val:   3,
		Left:  t5,
		Right: t6,
	}
	t2 := &TreeNode{
		Val:  2,
		Left: t4,
	}
	t1 := &TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	ans := findDuplicateSubtrees(t1)
	for _, a := range ans {
		fmt.Println(a.Val)
	}
}
