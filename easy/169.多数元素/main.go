package main

import "fmt"

// author: fengyuwusong date: 2022-08-11 16:35:41

// 169. 多数元素
//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,3]
//输出：3
//
// 示例 2：
//
//
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
//
//
//
//提示：
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
// Related Topics数组 | 哈希表 | 分治 | 计数 | 排序
//
// 👍 1523, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// 摩尔投票法
	// 只有一个数出现的次数大于数组长度的一半，那么这个数就是出现次数最多的数
	// 如果两个数出现的次数相等，那么这两个数就是出现次数最多的数
	// 如果两个数出现的次数不相等，那么这两个数就是出现次数最多的数
	count1 := 0
	count2 := 0
	num1 := 0
	num2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == num1 {
			count1++
		} else if nums[i] == num2 {
			count2++
		} else if count1 == 0 {
			num1 = nums[i]
			count1++
		} else if count2 == 0 {
			num2 = nums[i]
			count2++
		} else {
			count1--
			count2--
		}
	}
	if count1 > count2 {
		return num1
	}
	return num2
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{3, 2, 1, 1, 1, 5, 4}))
	fmt.Println(majorityElement([]int{6, 5, 5}))
}
