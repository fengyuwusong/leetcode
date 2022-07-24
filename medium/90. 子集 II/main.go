package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
//
//
//
// Related Topics位运算 | 数组 | 回溯
//
// 👍 882, 👎 0
//
//
//
//

// 2022-07-23 17:06:59 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	var backtrack func(start int, cur []int)
	backtrack = func(start int, cur []int) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			backtrack(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
}
