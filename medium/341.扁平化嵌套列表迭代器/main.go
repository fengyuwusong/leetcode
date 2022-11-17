package main

import "fmt"

// author: fengyuwusong date: 2022-08-19 16:00:59

// 341. 扁平化嵌套列表迭代器
//给你一个嵌套的整数列表 nestedList 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。请你实现一个迭代器将其扁平化
//，使之能够遍历这个列表中的所有整数。
//
// 实现扁平迭代器类 NestedIterator ：
//
//
// NestedIterator(List<NestedInteger> nestedList) 用嵌套列表 nestedList 初始化迭代器。
// int next() 返回嵌套列表的下一个整数。
// boolean hasNext() 如果仍然存在待迭代的整数，返回 true ；否则，返回 false 。
//
//
// 你的代码将会用下述伪代码检测：
//
//
//initialize iterator with nestedList
//res = []
//while iterator.hasNext()
//    append iterator.next() to the end of res
//return res
//
// 如果 res 与预期的扁平化列表匹配，那么你的代码将会被判为正确。
//
//
//
// 示例 1：
//
//
//输入：nestedList = [[1,1],2,[1,1]]
//输出：[1,1,2,1,1]
//解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
//
// 示例 2：
//
//
//输入：nestedList = [1,[4,[6]]]
//输出：[1,4,6]
//解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
//
//
//
//
// 提示：
//
//
// 1 <= nestedList.length <= 500
// 嵌套列表中的整数值在范围 [-10⁶, 10⁶] 内
//
// Related Topics栈 | 树 | 深度优先搜索 | 设计 | 队列 | 迭代器
//
// 👍 463, 👎 0
//
//
//
//

type NestedInteger struct {
	val    int
	list   []*NestedInteger
	isList bool
}

func (this NestedInteger) IsInteger() bool       { return !this.isList }
func (this NestedInteger) GetInteger() int       { return this.val }
func (this *NestedInteger) SetInteger(value int) { this.val = value; this.isList = false }
func (this *NestedInteger) Add(elem NestedInteger) {
	this.isList = true
	this.list = append(this.list, &elem)
}
func (this NestedInteger) GetList() []*NestedInteger { return this.list }

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: [][]*NestedInteger{nestedList}}
}

func (this *NestedIterator) Next() int {
	// 读栈顶
	queue := this.stack[len(this.stack)-1]
	// 出队
	cur := queue[0]
	this.stack[len(this.stack)-1] = queue[1:]
	return cur.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		// 出栈 获取队列
		queue := this.stack[len(this.stack)-1]
		// 当前队列空 出栈
		if len(queue) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		// 获取队列第一个元素
		curr := queue[0]
		if curr.IsInteger() {
			return true
		}
		// 否则为列表
		// 出队
		this.stack[len(this.stack)-1] = queue[1:]
		// 入栈
		this.stack = append(this.stack, curr.GetList())
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node1 := &NestedInteger{}
	node1.Add(NestedInteger{val: 1})
	node1.Add(NestedInteger{val: 1})
	node2 := &NestedInteger{val: 2}
	node3 := &NestedInteger{}
	node3.Add(NestedInteger{val: 1})
	node3.Add(NestedInteger{val: 1})
	nestedList := Constructor([]*NestedInteger{node1, node2, node3})
	for nestedList.HasNext() {
		fmt.Println(nestedList.Next())
	}
}
