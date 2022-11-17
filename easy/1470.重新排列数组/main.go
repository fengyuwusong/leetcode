package main

import "fmt"

// author: fengyuwusong date: 2022-08-29 14:16:16

// 1470. 重新排列数组
//给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。
//
// 请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。
//
//
//
// 示例 1：
//
// 输入：nums = [2,5,1,3,4,7], n = 3
//输出：[2,3,5,4,1,7]
//解释：由于 x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 ，所以答案为 [2,3,5,4,1,7]
//
//
// 示例 2：
//
// 输入：nums = [1,2,3,4,4,3,2,1], n = 4
//输出：[1,4,2,3,3,2,4,1]
//
//
// 示例 3：
//
// 输入：nums = [1,1,2,2], n = 2
//输出：[1,2,1,2]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 500
// nums.length == 2n
// 1 <= nums[i] <= 10^3
//
// Related Topics数组
//
// 👍 125, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func shuffle(nums []int, n int) []int {
	// 使用负号代表当前位置的元素已经正确 遍历每一个元素
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			j := i
			for nums[i] > 0 {
				// 计算j应该存放的位置
				if j < n {
					j *= 2
				} else {
					j = 2*(j%n) + 1
				}
				// 交换j i位置元素
				nums[j], nums[i] = nums[i], nums[j]
				// 此时j位置元素正确
				nums[j] = -nums[j]
				// i位置此时存放之前j位置的元素 进入下一次循环计算j位置元素应该存放的位置
			}
		}
	}
	// 将每一个元素转回正常
	for i := range nums {
		nums[i] = -nums[i]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}
