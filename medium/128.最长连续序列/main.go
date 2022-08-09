package main

import "fmt"

// author: fengyuwusong date: 2022-08-04 16:29:24

// 128. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
// Related Topics并查集 | 数组 | 哈希表
//
// 👍 1350, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	// 初始化并查集
	a := &ASCollection{
		parent:  make([]int, len(nums)),
		nodeNum: make([]int, len(nums)),
		hashMap: make(map[int]int),
	}
	// 初始化并查集
	for i := 0; i < len(nums); i++ {
		if _, ok := a.hashMap[nums[i]]; !ok {
			a.add(i, nums[i])
		}
	}
	return a.max()
}

// ASCollection 并查集实现
type ASCollection struct {
	parent  []int
	nodeNum []int
	hashMap map[int]int
}

func (a *ASCollection) add(index, value int) {
	a.parent[index] = index
	a.nodeNum[index] = 1
	a.hashMap[value] = index
	// 查找是否有相邻的值 加入集合
	if index2, ok := a.hashMap[value+1]; ok {
		a.union(index2, index)
	}
	if index2, ok := a.hashMap[value-1]; ok {
		// 大的加入小的集合
		a.union(index, index2)
	}
}

// 查找parent
func (a *ASCollection) find(x int) int {
	// x的父节点是自己时 说明是根节点
	if a.parent[x] == x {
		return x
	}
	// 路径压缩
	a.parent[x] = a.find(a.parent[x])
	return a.parent[x]
}

// x加入y的集合
func (a *ASCollection) union(x, y int) {
	xRoot := a.find(x)
	yRoot := a.find(y)
	a.parent[xRoot] = a.find(yRoot)
	a.nodeNum[yRoot] += a.nodeNum[xRoot]
}

// 返回最多子节点的数量
func (a *ASCollection) max() int {
	max := 0
	for _, v := range a.nodeNum {
		if v > max {
			max = v
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
