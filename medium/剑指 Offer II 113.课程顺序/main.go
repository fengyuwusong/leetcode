package main

import "fmt"

// author: fengyuwusong date: 2022-08-12 16:21:30

// 剑指 Offer II 113. 课程顺序
//现在总共有 numCourses 门课需要选，记为 0 到 numCourses-1。
//
// 给定一个数组 prerequisites ，它的每一个元素 prerequisites[i] 表示两门课程之间的先修顺序。 例如
//prerequisites[i] = [ai, bi] 表示想要学习课程 ai ，需要先完成课程 bi 。
//
// 请根据给出的总课程数 numCourses 和表示先修顺序的 prerequisites 得出一个可行的修课序列。
//
// 可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
//
//
//
// 示例 1:
//
//
//输入: numCourses = 2, prerequisites = [[1,0]]
//输出: [0,1]
//解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
//
// 示例 2:
//
//
//输入: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
//输出: [0,1,2,3] or [0,2,1,3]
//解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
// 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
//
//
// 示例 3:
//
//
//输入: numCourses = 1, prerequisites = []
//输出: [0]
//解释: 总共 1 门课，直接修第一门课就可。
//
//
//
// 提示:
//
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= numCourses * (numCourses - 1)
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// ai != bi
// prerequisites 中不存在重复元素
//
//
//
//
// 注意：本题与主站 210 题相同：https://leetcode-cn.com/problems/course-schedule-ii/
// Related Topics深度优先搜索 | 广度优先搜索 | 图 | 拓扑排序
//
// 👍 26, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
// dfs 后序遍历 翻转
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)
	onPath, isVisited := make([]bool, numCourses), make([]bool, numCourses)
	var ans []int
	var hasCycle bool
	var dfs func(curr int)
	dfs = func(curr int) {
		// 依靠之前的元素歩进到回来 存在环
		if onPath[curr] {
			hasCycle = true
			return
		}
		if isVisited[curr] || hasCycle {
			return
		}
		isVisited[curr] = true
		onPath[curr] = true
		// 遍历当前元素能到的下一个元素
		for _, v := range graph[curr] {
			dfs(v)
		}
		ans = append(ans, curr)
		onPath[curr] = false
	}
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	if hasCycle {
		return nil
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func buildGraph(courses int, prerequisites [][]int) [][]int {
	graph := make([][]int, courses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
}
