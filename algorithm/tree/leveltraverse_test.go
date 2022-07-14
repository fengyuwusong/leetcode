package tree

import (
	"container/list"
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输入一棵二叉树的根节点，层序遍历这棵二叉树
func levelTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	q := list.New()
	q.PushBack(root)

	// 从上到下遍历二叉树的每一层
	for q.Len() != 0 {
		sz := q.Len()
		// 从左到右遍历每一层的每个节点
		for i := 0; i < sz; i++ {
			ele := q.Front()
			q.Remove(ele)
			cur := ele.Value.(*TreeNode)
			res = append(res, cur.Val)
			// 将下一层节点放入队列
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
	}
	return res
}

func TestLevelTraverse(t *testing.T) {
	t5 := TreeNode{Val: 5}
	t4 := TreeNode{Val: 4}
	t2 := TreeNode{
		Val:   2,
		Left:  &t4,
		Right: &t5,
	}
	t3 := TreeNode{Val: 3}
	t1 := TreeNode{
		Val:   1,
		Left:  &t2,
		Right: &t3,
	}

	fmt.Println(levelTraverse(&t1))

}
