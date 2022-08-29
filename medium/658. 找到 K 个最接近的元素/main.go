package main

import (
	"fmt"
	"sort"
)

//给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//
// 整数 a 比整数 b 更接近 x 需要满足：
//
//
// |a - x| < |b - x| 或者
// |a - x| == |b - x| 且 a < b
//
//
//
//
// 示例 1：
//
//
//输入：arr = [1,2,3,4,5], k = 4, x = 3
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：arr = [1,2,3,4,5], k = 4, x = -1
//输出：[1,2,3,4]
//
//
//
//
// 提示：
//
//
// 1 <= k <= arr.length
// 1 <= arr.length <= 10⁴
// arr 按 升序 排列
// -10⁴ <= arr[i], x <= 10⁴
//
// Related Topics数组 | 双指针 | 二分查找 | 排序 | 堆（优先队列）
//
// 👍 358, 👎 0
//
//
//
//

// fengyuwusong 2022-08-25 01:02:59
//leetcode submit region begin(Prohibit modification and deletion)
func findClosestElements(arr []int, k int, x int) []int {
	// 二分法找到左右指针
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
			// 如果右指针 >= len(arr) || 左距离小于等于右距离
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}
