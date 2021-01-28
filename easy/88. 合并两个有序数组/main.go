package main

import "fmt"

/**
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。



示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]


提示：

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[i] <= 109
通过次数260,832提交次数529,912

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

// 由于两个数组有序，故从末尾开始比较大小可以得到最大值
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// 当p2>=0 意味nums2还有比nums1全体元素都小的元素需添加
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	// nums1 := []int{1}
	// nums2 := []int{}
	merge(nums1, 3, nums2, 3)
	// merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)
}
