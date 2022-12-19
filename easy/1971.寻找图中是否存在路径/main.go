package main

import "fmt"

//有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中
//edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
//
// 请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
//
// 给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回
//true，否则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
//输出：true
//解释：存在由顶点 0 到顶点 2 的路径:
//- 0 → 1 → 2
//- 0 → 2
//
//
// 示例 2：
//
//
//输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
//
//输出：false
//解释：不存在由顶点 0 到顶点 5 的路径.
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2 * 10⁵
// 0 <= edges.length <= 2 * 10⁵
// edges[i].length == 2
// 0 <= ui, vi <= n - 1
// ui != vi
// 0 <= source, destination <= n - 1
// 不存在重复边
// 不存在指向顶点自身的边
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 66 👎 0

// author: fengyuwusong
// create time: 2022-12-19 10:12:00
// leetcode submit region begin(Prohibit modification and deletion)
func validPath(n int, edges [][]int, source int, destination int) bool {
	// bfs
	queue := []int{source}
	visited := make(map[int]bool, n)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true
		if u == destination {
			return true
		}
		for _, edge := range edges {
			if edge[0] == u && !visited[edge[1]] {
				queue = append(queue, edge[1])
				continue
			}
			if edge[1] == u && !visited[edge[0]] {
				queue = append(queue, edge[0])
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
	fmt.Println(validPath(6, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5))
}
