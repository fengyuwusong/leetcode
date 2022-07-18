package easy

import (
	"fmt"
	"testing"
)

/**
不使用任何内建的哈希表库设计一个哈希集合

具体地说，你的设计应该包含以下的功能

add(value)：向哈希集合中插入一个值。
contains(value) ：返回哈希集合中是否存在这个值。
remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

示例:

MyHashSet hashSet = new MyHashSet();
hashSet.add(1);
hashSet.add(2);
hashSet.contains(1);    // 返回 true
hashSet.contains(3);    // 返回 false (未找到)
hashSet.add(2);
hashSet.contains(2);    // 返回 true
hashSet.remove(2);
hashSet.contains(2);    // 返回  false (已经被删除)

注意：

所有的值都在 [0, 1000000]的范围内。
操作的总数目在[1, 10000]范围内。
不要使用内建的哈希集合库。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-hashset
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 单链表实现
// 时间复杂度 O(N/K) n是指可能值数量 K是桶数  假设值平均分布 则每个桶的平均大小是 N/K 那么每个操作最坏情况需扫描整个桶 因此时间复杂度O(N/K)
// 空间复杂度 O(K+N)
// 解法2 使用二叉搜索树作为桶 时间复杂度O(logN/K) 二分查找 空间复杂度O(K+N)

// 解法2 todo 二叉树实现
type ListNode struct {
	Val  int
	Next *ListNode
}
type MyHashSet struct {
	val []*ListNode
}

func hashFun(x int) int {
	// 使用较大的质数可减少碰撞
	return x % 768
}

/** Initialize your data structure here. */
func ConstructorMyHashSet() MyHashSet {
	return MyHashSet{
		val: make([]*ListNode, 768),
	}
}

func (this *MyHashSet) Add(key int) {
	n := hashFun(key)
	cur := this.val[n]
	if cur == nil {
		this.val[n] = &ListNode{Val: key}
		return
	}
	if cur.Val == key {
		return
	}
	pre := cur
	cur = cur.Next
	for cur != nil {
		if cur.Val == key {
			return
		}
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = &ListNode{Val: key}
}

func (this *MyHashSet) Remove(key int) {
	n := hashFun(key)
	cur := this.val[n]
	if cur == nil {
		return
	}
	if cur.Val == key {
		this.val[n] = cur.Next
		return
	}
	pre := cur
	cur = cur.Next
	for cur != nil {
		if cur.Val == key {
			pre.Next = cur.Next
			return
		}
		cur = cur.Next
		pre = pre.Next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	n := hashFun(key)
	cur := this.val[n]
	if cur == nil {
		return false
	}
	for cur != nil {
		if cur.Val == key {
			return true
		}
		cur = cur.Next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func TestMyHashSet(t *testing.T) {
	mySet := ConstructorMyHashSet()
	mySet.Add(1)
	mySet.Add(2)
	fmt.Println(mySet.Contains(2))
	fmt.Println(mySet.Contains(3))
	mySet.Add(12)
	mySet.Remove(2)
	mySet.Add(2)
	mySet.Add(12)
	mySet.Add(22)
	fmt.Println(mySet.Contains(12))
	mySet.Remove(12)
	fmt.Println(mySet.Contains(12))
	fmt.Println(mySet.Contains(22))
	fmt.Println(mySet.Contains(2))
}
