package main

import (
	"fmt"
	"sort"
)

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
// 示例 2:
//
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
// Related Topics数组 | 回溯
//
// 👍 1041, 👎 0
//
//
//
//

// 2022-07-24 16:20:11 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var result [][]int
	var backtrack func(start int, curInts []int, cur int)
	backtrack = func(start int, curInts []int, cur int) {
		if cur == target {
			temp := make([]int, len(curInts))
			copy(temp, curInts)
			result = append(result, temp)
			return
		}
		if cur > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			// 剪枝
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			if candidates[i] > target {
				continue
			}
			if cur+candidates[i] > target {
				continue
			}
			// 选择
			curInts = append(curInts, candidates[i])
			backtrack(i+1, curInts, cur+candidates[i])
			// 回溯
			curInts = curInts[:len(curInts)-1]
		}
	}
	backtrack(0, []int{}, 0)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(combinationSum2([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum2([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
