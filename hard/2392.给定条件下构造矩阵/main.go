package main

import "fmt"

// author: fengyuwusong date: 2022-09-05 16:33:59

// 2392. 给定条件下构造矩阵
//给你一个 正 整数 k ，同时给你：
//
//
// 一个大小为 n 的二维整数数组 rowConditions ，其中 rowConditions[i] = [abovei, belowi] 和
// 一个大小为 m 的二维整数数组 colConditions ，其中 colConditions[i] = [lefti, righti] 。
//
//
// 两个数组里的整数都是 1 到 k 之间的数字。
//
// 你需要构造一个 k x k 的矩阵，1 到 k 每个数字需要 恰好出现一次 。剩余的数字都是 0 。
//
// 矩阵还需要满足以下条件：
//
//
// 对于所有 0 到 n - 1 之间的下标 i ，数字 abovei 所在的 行 必须在数字 belowi 所在行的上面。
// 对于所有 0 到 m - 1 之间的下标 i ，数字 lefti 所在的 列 必须在数字 righti 所在列的左边。
//
//
// 返回满足上述要求的 任意 矩阵。如果不存在答案，返回一个空的矩阵。
//
//
//
// 示例 1：
//
//
//
// 输入：k = 3, rowConditions = [[1,2],[3,2]], colConditions = [[2,1],[3,2]]
//输出：[[3,0,0],[0,0,1],[0,2,0]]
//解释：上图为一个符合所有条件的矩阵。
//行要求如下：
//- 数字 1 在第 1 行，数字 2 在第 2 行，1 在 2 的上面。
//- 数字 3 在第 0 行，数字 2 在第 2 行，3 在 2 的上面。
//列要求如下：
//- 数字 2 在第 1 列，数字 1 在第 2 列，2 在 1 的左边。
//- 数字 3 在第 0 列，数字 2 在第 1 列，3 在 2 的左边。
//注意，可能有多种正确的答案。
//
//
// 示例 2：
//
// 输入：k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
//输出：[]
//解释：由前两个条件可以得到 3 在 1 的下面，但第三个条件是 3 在 1 的上面。
//没有符合条件的矩阵存在，所以我们返回空矩阵。
//
//
//
//
// 提示：
//
//
// 2 <= k <= 400
// 1 <= rowConditions.length, colConditions.length <= 10⁴
// rowConditions[i].length == colConditions[i].length == 2
// 1 <= abovei, belowi, lefti, righti <= k
// abovei != belowi
// lefti != righti
//
// Related Topics图 | 拓扑排序 | 数组 | 矩阵
//
// 👍 24, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowGraph := buildGraph(k, rowConditions)
	colGraph := buildGraph(k, colConditions)
	isVisited := make([]bool, k+1)
	onPath := make([]bool, k+1)
	// 行列的拓扑排序
	var rows []int
	var cols []int
	hasCycle := false
	var dfs func(graph [][]int, ans *[]int, x int)
	dfs = func(graph [][]int, ans *[]int, x int) {
		if onPath[x] {
			hasCycle = true
			return
		}
		if isVisited[x] || hasCycle {
			return
		}
		isVisited[x] = true
		onPath[x] = true
		for _, y := range graph[x] {
			dfs(graph, ans, y)
		}
		*ans = append(*ans, x)
		onPath[x] = false
	}
	for i := 1; i <= k; i++ {
		dfs(rowGraph, &rows, i)
	}
	if hasCycle {
		return nil
	}
	isVisited = make([]bool, k+1)
	onPath = make([]bool, k+1)
	for i := 1; i <= k; i++ {
		dfs(colGraph, &cols, i)
	}
	if hasCycle {
		return nil
	}

	local := make(map[int][]int)
	for i := 0; i < len(rows); i++ {
		local[rows[i]] = []int{i, 0}
	}
	for i := 0; i < len(cols); i++ {
		if val, ok := local[cols[i]]; ok {
			val[1] = i
		} else {
			local[cols[i]] = []int{0, i}
		}
	}
	ans := make([][]int, k)
	for i := range ans {
		ans[i] = make([]int, k)
	}
	for k, v := range local {
		i := v[0]
		j := v[1]
		if i == 0 {
			for ; ans[i][j] != 0 && i < k; i++ {
			}
			if i == k {
				return nil
			}
		}
		if j == 0 {
			for ; ans[i][j] != 0 && j < k; j++ {
			}
			if j == k {
				return nil
			}
		}
		ans[i][j] = k
	}
	return ans
}

func buildGraph(k int, conditions [][]int) [][]int {
	graph := make([][]int, k+1)
	for _, condition := range conditions {
		graph[condition[1]] = append(graph[condition[1]], condition[0])
	}
	return graph
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	matrix := buildMatrix(3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}})
	fmt.Println(matrix)
	matrix = buildMatrix(4,
		[][]int{{4, 3}, {1, 4}, {1, 2}, {4, 3}, {4, 3}},         // 1-4-2-3 1-2-4-3
		[][]int{{3, 2}, {1, 4}, {2, 1}, {3, 4}, {3, 1}, {2, 4}}) // 3-2-1-4
	fmt.Println(matrix)
}
