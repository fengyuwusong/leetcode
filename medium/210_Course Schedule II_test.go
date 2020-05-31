package medium

import (
	"fmt"
	"testing"
)

// dfs 深度优先遍历 时间复杂度 O(n+m) 空间复杂度 O(n+m) n: 课程数 m: 先修课程的要求数
//func findOrder(numCourses int, prerequisites [][]int) []int {
//	var (
//		cs      = make([][]int, numCourses) // 记录向量图 c[课程x] => [y, z] 学习课程x后可以学习的课程
//		visited = make([]int, numCourses)   // 记录是否搜索过 0 => 未搜索 1 => 搜索中 2 => 已完成搜索
//		res     []int                       // 结果
//		invalid bool                        // 是否合法
//		dfs     func(cNum int)              // dfs 算法
//	)
//	// 定义dfs算法
//	dfs = func(cNum int) {
//		visited[cNum] = 1
//		// 查找可以学习的课程
//		for _, ccNum := range cs[cNum] {
//			if visited[ccNum] == 0 {
//				dfs(ccNum)
//			} else if visited[ccNum] == 1 {
//				invalid = true
//				break
//			}
//		}
//		visited[cNum] = 2
//		res = append(res, cNum)
//	}
//
//	// 记录向量图
//	for _, p := range prerequisites {
//		cs[p[1]] = append(cs[p[1]], p[0])
//	}
//
//	// 遍历每个课程
//	for i := 0; i < numCourses && !invalid; i++ {
//		if visited[i] == 0 {
//			dfs(i)
//		}
//	}
//	// 非法返回空
//	if invalid {
//		return nil
//	}
//
//	// 反转数组
//	for i := 0; i < numCourses/2; i++ {
//		res[i], res[numCourses-i-1] = res[numCourses-i-1], res[i]
//	}
//
//	return res
//}

// bfs 广度优先遍历
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		cs    = make([][]int, numCourses) // 记录向量图 c[课程x] => [y, z] 学习课程x之前必须学习的课程 即入度
		indeg = make([]int, numCourses)   // 记录节点入度状态
		res   []int                       // 结果
	)

	// 构造向量图
	for _, p := range prerequisites {
		cs[p[1]] = append(cs[p[1]], p[0])
		indeg[p[0]]++
	}

	// 遍历每个课程 获取入度值为0的节点
	var topNode []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			topNode = append(topNode, i)
		}
	}
	// 从顶部入度值为0的节点开始遍历并摘除
	for i := 0; i < len(topNode); i++ {
		// 遍历学了cNum之后可以学习的课程 入度-1
		for _, ccNum := range cs[topNode[i]] {
			indeg[ccNum]--
			// 入度为0为顶部节点
			if indeg[ccNum] == 0 {
				topNode = append(topNode, ccNum)
			}
		}
		// 加入结果
		res = append(res, topNode[i])
	}

	// 结果 != 全部课程数量说明存在闭环 不存在拓扑排序
	if len(res) != numCourses {
		return nil
	}

	return res
}

func TestFindOrder(t *testing.T) {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
