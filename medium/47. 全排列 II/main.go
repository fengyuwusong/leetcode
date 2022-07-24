package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics数组 | 回溯
//
// 👍 1137, 👎 0
//
//
//
//

// 2022-07-24 16:34:42 fengyuwusong
//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	// 回溯
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func([]int, []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 剪枝逻辑 保证相同值元素相对顺序一致 当前相同值必须在前面相同的元素使用之后才能使用
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(path, used)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}
