package main

import "fmt"

// author: fengyuwusong date: 2022-08-11 17:09:02

// 207. 课程表
//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//
// 提示：
//
//
// 1 <= numCourses <= 10⁵
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
//
// Related Topics深度优先搜索 | 广度优先搜索 | 图 | 拓扑排序
//
// 👍 1377, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 建立图
	graph := buildGraph(numCourses, prerequisites)
	// 判断是否已访问 用于剪枝
	isVisited := make([]bool, numCourses)
	// 记录访问路径 用于判断
	onPath := make([]bool, numCourses)
	// 判断是否有环
	hasCycle := false
	// 判断从x号课程开始是否可以完成所有课程
	var dfs func(x int)
	dfs = func(x int) {
		// 存在环
		if onPath[x] {
			hasCycle = true
			return
		}
		// 已访问过 返回
		if isVisited[x] || hasCycle {
			return
		}
		// x节点已访问
		isVisited[x] = true
		onPath[x] = true
		for _, y := range graph[x] {
			dfs(y)
		}
		onPath[x] = false
	}
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	return !hasCycle
}

func buildGraph(courses int, prerequisites [][]int) [][]int {
	graph := make([][]int, courses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(3, [][]int{{0, 1}, {1, 2}}))
}
