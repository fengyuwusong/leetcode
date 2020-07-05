package medium

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 根据快排的思路 可以随机取一个数 得到其基准位置 再二分缩小范围逐渐得到第K大元素 时间复杂度O(n) 空间复杂度O(logn)
// 解法2 基于堆排序的选择方法 时间复杂度O(n+klogn)=>O(nlogn) 空间复杂度O(logn)

// 解法1
func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	// 转换一下，第 k 大元素的索引是 len - k
	target := len(nums) - k
	l, r := 0, len(nums)-1
	for i := getNumsKth(nums, l, r); target != i; i = getNumsKth(nums, l, r) {
		// 大于查找左区间
		if i > target {
			r = i - 1
		} else if i < target {
			//	小于查找右区间
			l = i + 1
		}
	}
	return nums[target]
}

// 获取nums数组中任一元素对应在排序数组后的索引位置
func getNumsKth(nums []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(r-l+1) + l
	// 将其置于最后一位
	nums[r], nums[pivotIndex] = nums[pivotIndex], nums[r]
	// 双指针遍历 将大于 nums[r]的数置于nums[i]中
	i, j := l-1, l
	for ; j < r; j++ {
		// 小于则与++i替换位置
		if nums[j] < nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 将nums[r]与 nums[i+1]交换 即 [l, i]区间数均小于 nums[i] [i+2, r]区间数均大于nums[i+1]
	nums[i+1], nums[r] = nums[r], nums[i+1]
	// 返回当前元素pivotIndex 实际在排序后数组的索引
	return i + 1
}

// 解法2 堆排序确定位置
func findKthLargest2(nums []int, k int) int {
	// 构造大顶堆 从第一个的非叶子节点开始调整顺序以符合大顶堆定义
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	// 由大顶堆定义 目前根节点为最大数 已确认第一大的数 交换并重建确认后续的数
	// 当j=len(nums)-k 时 说明[j, len(nums)-1] 区间已被排序 即可退出循环取出第一个节点即为第k大的数=> 索引 len(num)-k=j
	for j := len(nums) - 1; j >= len(nums)-k+1; j-- {
		// 交换位置
		nums[j], nums[0] = nums[0], nums[j]
		adjustHeap(nums, 0, j)
	}
	return nums[0]
}

func adjustHeap(nums []int, i, length int) {
	temp := nums[i]
	// 从当前节点往下遍历所有左右叶子节点 若大于它则交换位置
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && nums[k] < nums[k+1] {
			k++
		}
		if temp < nums[k] {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}
	nums[i] = temp
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{1, 3, 2, 6, 2}, 5))
	fmt.Println(findKthLargest([]int{1, 3}, 1))
	fmt.Println(findKthLargest([]int{1}, 1))
	fmt.Println(findKthLargest2([]int{1, 3, 2, 6, 2}, 5))
	fmt.Println(findKthLargest2([]int{1, 3}, 1))
	fmt.Println(findKthLargest2([]int{1}, 1))
}
