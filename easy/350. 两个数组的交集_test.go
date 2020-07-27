package easy

import (
	"fmt"
	"testing"
)

/**
给定两个数组，编写一个函数来计算它们的交集。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 解法1 排序后遍历 时间复杂度O(mlogm+nlogn) 空间复杂度O(min(m,n))
// 解法2 hash O(m+n) 空间复杂度O(min(m,n))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	hashMap := make(map[int]int)
	var ans []int
	for i := range nums1 {
		hashMap[nums1[i]]++
	}
	for i := range nums2 {
		if hashMap[nums2[i]] > 0 {
			ans = append(ans, nums2[i])
			hashMap[nums2[i]]--
		}
	}
	return ans
}

func TestIntersect(t *testing.T) {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 5, 9}, []int{9, 4, 9, 8, 4}))
}
