package main

import "fmt"

/**
704. 二分查找
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1


提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
 */

func search(nums []int, target int) int {
    // 获取中间值
    minIndex := 0
    maxIndex := len(nums) - 1
    midIndex := maxIndex/2
    for minIndex <= maxIndex {
        // 结束递归条件
        if nums[midIndex] == target {
            return midIndex
        }

        if nums[midIndex] > target {
            maxIndex = midIndex - 1
        }

        if nums[midIndex] < target {
            minIndex = midIndex + 1
        }
        midIndex = (maxIndex + minIndex) / 2
    }
    return -1
}


func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 5))
}
