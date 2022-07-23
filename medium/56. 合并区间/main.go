package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		// 重复区间
		if intervals[i][0] >= res[len(res)-1][0] && intervals[i][1] > res[len(res)-1][1] && intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = intervals[i][1]
		} else if intervals[i][0] > res[len(res)-1][0] && intervals[i][1] > res[len(res)-1][1] && intervals[i][0] > res[len(res)-1][1] {
			// 非重复区间
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
