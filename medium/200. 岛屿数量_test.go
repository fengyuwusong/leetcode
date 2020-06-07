package medium

import (
	"fmt"
	"testing"
)

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1:
输入:
11110
11010
11000
00000
输出: 1

示例 2:
输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
*/

// 解法1：dfs 时间复杂度O(nm) 空间复杂度 O(n+m) dfs深度
// 解法2：使用bfs广度优先遍历 当点为1时扩散范围(上下左右)并标记岛屿编号
// 时间复杂度O(nm) 空间复杂度O(min(n,m)) bfs队列长度为矩阵对角线 全是陆地时
// 将将岛屿上下位移往下位移则可不用考虑边界值问题 但会损耗空间(程序输出无问题 leetcode输出错误)
// 解法3：todo 并查集 扫描二维网格 如果一个位置为1 则将其右下方向为1在并查集中合并 最终岛屿的数量就是并查集中联通分量的数目
// 时间复杂度O(nm * a(nm))注意当使用路径压缩（见 find 函数）和按秩合并（见数组 rank）实现并查集时，
// 单次操作的时间复杂度为 a(MN)α(MN)，其中 α(x) 为反阿克曼函数，
// 当自变量 x 的值在人类可观测的范围内（宇宙中粒子的数量）时，函数 α(x) 的值不会超过 55，因此也可以看成是常数时间复杂度。
// 空间复杂度O(nm)

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int

//bfs
func numIslands(grid [][]byte) int {
	row = len(grid)
	if row == 0 {
		return 0
	}
	col = len(grid[0])
	// 岛屿个数
	ans := 0
	// 遍历grid
	for i, v := range grid {
		for j, vv := range v {
			// 是陆地
			if vv == '1' {
				bfs(grid, i, j)
				ans++
			}
		}
	}
	return ans
}

// dfs
func numIslandsDfs(grid [][]byte) int {
	row = len(grid)
	if row == 0 {
		return 0
	}
	col = len(grid[0])
	// 岛屿个数
	ans := 0
	// 遍历grid
	for i, v := range grid {
		for j, vv := range v {
			// 是陆地
			if vv == '1' {
				dfs(grid, i, j)
				ans++
			}
		}
	}
	return ans
}

type UnionFind struct {
	count  int   // 联通数量
	rank   []int // 层级
	parent []int // 上一级
}

// 并查集初始化方法
func (uf *UnionFind) Init(grid [][]byte) {
	uf.count = 0
	m := len(grid)
	n := len(grid[0])
	uf.parent = make([]int, n*m)
	uf.rank = make([]int, n*m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 将其上级设为自己
				uf.parent[i*n+j] = i*n + j
				// 节点数量+1
				uf.count++
			}
			// 当前层级置0
			uf.rank[i*n+j] = 0
		}
	}
}

// 路径查找方法
func (uf *UnionFind) find(i int) int {
	// 上级并不等与自己
	if uf.parent[i] != i {
		// 继续查找其下级 压缩路径
		uf.parent[i] = uf.find(uf.parent[i])
	}
	return uf.parent[i]
}

// 合并集合
func (uf *UnionFind) union(x, y int) {
	// 获取集合x y的最上级集合
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX] += 1
		}
		uf.count -= 1
	}
}

// 并查集
func numIslandsUnionFindSet(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	// 初始化并查集
	unionFind := &UnionFind{}
	unionFind.Init(grid)
	// 遍历数组 当前节点时且下、右节点是陆地的进行合并
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// 已遍历将此处变为水
				grid[i][j] = 0
				for m := 0; m < 4; m++ {
					ii := i + dx[m]
					jj := j + dy[m]
					if ii >= 0 && ii < row && jj >= 0 && jj < col && grid[ii][jj] == '1' {
						// 将当前节点与四周陆地节点合并
						unionFind.union(i*col+j, ii*col+jj)
					}
				}
			}
		}
	}
	return unionFind.count
}

func bfs(grid [][]byte, i, j int) {
	queue := make([]int, len(grid)*2)
	queue = []int{i, j}
	// 遍历过则将此区域变为水
	grid[i][j] = 0
	for len(queue) != 0 {
		i, j := queue[0], queue[1]
		// 将头节点移出队列
		queue = queue[2:]
		// 遍历上下左右节点
		for m := 0; m < 4; m++ {
			ii := i + dx[m]
			jj := j + dy[m]
			if ii >= 0 && ii < row && jj >= 0 && jj < col && grid[ii][jj] == '1' {
				queue = append(queue, ii, jj)
				// 已遍历将此处变为水
				grid[ii][jj] = 0
			}
		}
	}
}

func dfs(grid [][]byte, i, j int) {
	// 遍历过则将此区域变为水
	grid[i][j] = '0'
	// 对上下左右大陆节点进行dfs
	for m := 0; m < 4; m++ {
		ii := i + dx[m]
		jj := j + dy[m]
		if ii >= 0 && ii < row && jj >= 0 && jj < col && grid[ii][jj] == '1' {
			dfs(grid, ii, jj)
		}
	}
}

func TestNumIsLands(t *testing.T) {
	fmt.Println(numIslandsUnionFindSet([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslandsUnionFindSet([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}
