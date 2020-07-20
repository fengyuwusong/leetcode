package medium

/**
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的 深拷贝。

我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。


示例 1：



输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：



输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：



输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
示例 4：

输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。


提示：

-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for a Node.
type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

// 解法1: hash表记录当前节点 在重复遍历一次写入新的random节点 时间复杂度O(n) 空间复杂度O(n)
// 解法2: 回溯法 将链表视为图
// 将遇到过的节点放到字典中 如已在字典中的节点则不再访问 没在字典中的节点创建并赋值 即可完成深度遍历
// 时间空间复杂度同解法1
// 解法3: 拷贝法
// 第一次遍历链表拷贝每个节点并将新的节点赋值位于原节点的next
// 第二次遍历赋值偶数节点的random, 其值等于原节点的random的next节点
// 第三次遍历去除奇数节点
// 时间复杂度 O(n) 空间复杂度O(1) 结果链表不视为暂用空间

// 解法1
func copyRandomList(head *Node138) *Node138 {
	if head == nil {
		return nil
	}

	hashMap := make(map[*Node138]*Node138)
	root := &Node138{}
	cur := root
	for head != nil {
		n := &Node138{Val: head.Val, Random: head.Random}
		cur.Next = n
		cur = n
		hashMap[head] = n
		head = head.Next
	}
	cur = root.Next
	for cur != nil {
		if cur.Random != nil {
			cur.Random = hashMap[cur.Random]
		}
		cur = cur.Next
	}
	return root.Next
}
