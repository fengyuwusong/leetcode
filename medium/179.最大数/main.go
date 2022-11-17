package main

import (
	"fmt"
	"sort"
	"strconv"
)

// author: fengyuwusong date: 2022-08-19 11:58:15

// 179. 最大数
//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
//
//输入：nums = [10,2]
//输出："210"
//
// 示例 2：
//
//
//输入：nums = [3,30,34,5,9]
//输出："9534330"
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 10⁹
//
// Related Topics贪心 | 字符串 | 排序
//
// 👍 992, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i int, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a+b > b+a
	})
	var ans string
	for i := 0; i < len(nums); i++ {
		ans += fmt.Sprintf("%d", nums[i])
	}
	if ans[0] == '0' {
		return "0"
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 01, 02, 03}))
	fmt.Println(largestNumber([]int{0, 0}))
}
