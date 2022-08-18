package main

import "fmt"

// author: fengyuwusong date: 2022-08-12 10:37:18

// 210. 课程表 II
//现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
//prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
//
//
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
//
//
// 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
//
//
//
// 示例 1：
//
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：[0,1]
//解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
//
//
// 示例 2：
//
//
//输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
//输出：[0,2,1,3]
//解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
//因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
//
// 示例 3：
//
//
//输入：numCourses = 1, prerequisites = []
//输出：[0]
//
//
//
//提示：
//
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= numCourses * (numCourses - 1)
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// ai != bi
// 所有[ai, bi] 互不相同
//
// Related Topics深度优先搜索 | 广度优先搜索 | 图 | 拓扑排序
//
// 👍 680, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
// bfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)
	inNums := make([]int, numCourses)
	// 计算入度
	for _, to := range graph {
		for _, v := range to {
			inNums[v]++
		}
	}
	var queue []int
	// 从入度为0的进入队列
	for i, v := range inNums {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	var ans []int
	for len(queue) != 0 {
		// 出队
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		ans = append(ans, curr)
		// 将当前出队元素可达的元素入度-1 入度=0说明可学习 加入队列
		for _, v := range graph[curr] {
			inNums[v]--
			if inNums[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 存在环的话 len(ans) != numCourses 因为互相依赖, 环元素的入度!=0
	if len(ans) != numCourses {
		return nil
	}
	return ans
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 1}, {1, 2}, {2, 3}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
}
