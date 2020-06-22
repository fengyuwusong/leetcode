package medium

import (
	"fmt"
	"sort"
	"testing"
)

/**
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 贪心法
// 使用merge存储答案
// 数组按照区间左端点排序 遍历数组 当任一区间左端端点大于数组merge中最后一个区间的右端点时, 说明两区间不重合 直接将当前区间加入merge末尾
// 反之则重合 合并当前区间与merge最后一个区间 左端点使用merge最后一个区间的左端点 右端点使用当前区间右端点与merge最后区间右端点的最大值
func merge(intervals [][]int) [][]int {
	intervalsLen := len(intervals)
	if intervalsLen <= 0 {
		return nil
	}
	// 根据区间左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 初始化 将第一个区间加入merge
	merge := [][]int{intervals[0]}
	for i := 1; i < intervalsLen; i++ {
		// 区间不重复
		if intervals[i][0] > merge[len(merge)-1][1] {
			merge = append(merge, intervals[i])
		}
		// 区间重复 合并区间 更新merge最后一个区间的右端点
		if intervals[i][1] > merge[len(merge)-1][1] {
			merge[len(merge)-1][1] = intervals[i][1]
		}
	}
	return merge
}

func TestMerge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
