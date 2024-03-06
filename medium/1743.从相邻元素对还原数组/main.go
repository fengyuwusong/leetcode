package main

import (
	"fmt"
	"math"
)

// 存在一个由 n 个不同元素组成的整数数组 nums ，但你已经记不清具体内容。好在你还记得 nums 中的每一对相邻元素。
//
// 给你一个二维整数数组 adjacentPairs ，大小为 n - 1 ，其中每个 adjacentPairs[i] = [ui, vi] 表示元素
// ui 和 vi 在 nums 中相邻。
//
// 题目数据保证所有由元素 nums[i] 和 nums[i+1] 组成的相邻元素对都存在于 adjacentPairs 中，存在形式可能是 [nums[i]
// , nums[i+1]] ，也可能是 [nums[i+1], nums[i]] 。这些相邻元素对可以 按任意顺序 出现。
//
// 返回 原始数组 nums 。如果存在多种解答，返回 其中任意一个 即可。
//
//
//
// 示例 1：
//
//
// 输入：adjacentPairs = [[2,1],[3,4],[3,2]]
// 输出：[1,2,3,4]
// 解释：数组的所有相邻元素对都在 adjacentPairs 中。
// 特别要注意的是，adjacentPairs[i] 只表示两个元素相邻，并不保证其 左-右 顺序。
//
//
// 示例 2：
//
//
// 输入：adjacentPairs = [[4,-2],[1,4],[-3,1]]
// 输出：[-2,4,1,-3]
// 解释：数组中可能存在负数。
// 另一种解答是 [-3,1,4,-2] ，也会被视作正确答案。
//
//
// 示例 3：
//
//
// 输入：adjacentPairs = [[100000,-100000]]
// 输出：[100000,-100000]
//
//
//
//
// 提示：
//
//
// nums.length == n
// adjacentPairs.length == n - 1
// adjacentPairs[i].length == 2
// 2 <= n <= 10⁵
// -10⁵ <= nums[i], ui, vi <= 10⁵
// 题目数据保证存在一些以 adjacentPairs 作为元素对的数组 nums
//
//
// Related Topics 数组 哈希表 👍 122 👎 0

// author: fengyuwusong
// create time: 2024-03-06 18:30:18
// leetcode submit region begin(Prohibit modification and deletion)
func restoreArray(adjacentPairs [][]int) []int {
	// 构造 hash 表记录每个数字的左右
	hash := make(map[int][]int)
	for _, pair := range adjacentPairs {
		hash[pair[0]] = append(hash[pair[0]], pair[1])
		hash[pair[1]] = append(hash[pair[1]], pair[0])
	}

	// 找出仅有一个相邻的数，则为左右边界
	start := math.MaxInt32
	for k, v := range hash {
		if len(v) == 1 {
			start = k
		}
	}
	if start == math.MaxInt32 {
		return nil
	}
	// 从左边界开始诸葛插入
	res := []int{start, hash[start][0]}
	for {
		lastNum := res[len(res)-2]
		start = res[len(res)-1]
		if len(hash[start]) == 1 {
			break
		}

		for _, v := range hash[start] {
			if v != lastNum {
				res = append(res, v)
			}
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}}))
}
