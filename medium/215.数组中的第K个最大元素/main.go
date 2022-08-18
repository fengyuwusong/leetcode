package main

import (
	"fmt"
	"math/rand"
	"time"
)

// author: fengyuwusong date: 2022-08-15 17:08:58

// 215. 数组中的第K个最大元素
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
// Related Topics数组 | 分治 | 快速选择 | 排序 | 堆（优先队列）
//
// 👍 1810, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
// 快排分治
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right int, k int) int {
	r := rand.Intn(right-left+1) + left
	nums[r], nums[right] = nums[right], nums[r]
	// 查找数组最后一个对应的位置
	index := findIndex(nums, left, right)
	// 计算第几大
	maxNum := len(nums) - index
	if maxNum == k {
		return nums[index]
	} else if maxNum < k {
		return quickSelect(nums, left, index-1, k)
	} else if maxNum > k {
		return quickSelect(nums, index+1, right, k)
	}
	return -1
}

func findIndex(nums []int, left, right int) int {
	i, j := left, right
	temp := nums[right]
	for i < j {
		for i < j && nums[i] < temp {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
		for i < j && nums[j] > temp {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
	}
	nums[i] = temp
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 1))
	fmt.Println(findKthLargest([]int{2, 1}, 1))
	fmt.Println(findKthLargest([]int{3, 3, 3, 3, 3, 3, 3, 3, 3}, 1))
	fmt.Println(findKthLargest([]int{-1, 2, 0}, 1))
}
