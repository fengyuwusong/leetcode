package sort

import (
	"fmt"
	"testing"
)

// 快速排序
// 第某一元素nums[x]为基准 双指针分别头尾遍历剩余元素
// 当遇到 num[i]>nums[x] 且nums[j]<nums[x] 时两数交换位置
// 当i=j时 x 和i替换位置 可将数组分为三部分 小于nums[x] nums[x] 大于nums[x]
// 根据分治的方法重复此过程则可将数组排序
// 时间复杂度O(nlogn) 空间复杂度O(logn) 递归栈
// 时间复杂度坏情况 O(n^2) 当每次选出中值为最值时, 即数组已顺序或逆序时 则需要每个元素都进行划分
// 解决办法: 1.随机取值 2. 三点中值
// 当数组规模较小时 插入排序比快排时间快

// 参考
// https://www.cnblogs.com/nicaicai/p/12689403.html
// https://blog.csdn.net/qq_29110265/article/details/83502450

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, startIndex, endIndex int) {
	// 三点中值取基准数
	pivot := geMedian(nums, startIndex, endIndex)
	if endIndex-startIndex+1 >= 3 {
		// 定义双指针 由于getMedian前中后已排序 倒数第二为pivot 故i j从 startIndex+1 endIndex-2开始
		i, j := startIndex+1, endIndex-2
		for {
			// 分别找到指针中两者大于pivot 和小于pivot的数
			for ; nums[i] < pivot; i++ {
			}
			for ; nums[j] > pivot; j-- {
			}
			// 找到且两者未相遇 替换
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				// 相遇 跳出循环
				break
			}
		}
		// 将pivot值置于i,j找到的位置
		nums[i], nums[endIndex-1] = nums[endIndex-1], nums[i]
		// 分治左右半区
		quickSort(nums, startIndex, i-1)
		quickSort(nums, i+1, endIndex)
	}
}

func geMedian(nums []int, startIndex, endIndex int) int {
	if endIndex-startIndex+1 >= 3 {
		centerIndex := (endIndex + startIndex) / 2
		// 将左中右三个数进行排序交换
		if nums[startIndex] > nums[endIndex] {
			nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
		}
		if nums[startIndex] > nums[centerIndex] {
			nums[startIndex], nums[centerIndex] = nums[centerIndex], nums[startIndex]
		}
		if nums[centerIndex] > nums[endIndex] {
			nums[endIndex], nums[centerIndex] = nums[centerIndex], nums[endIndex]
		}
		// 将中间值至于右边第二个数, 确保右边扫描的数都大于中间值 方便排序时插入
		nums[centerIndex], nums[endIndex-1] = nums[endIndex-1], nums[centerIndex]
		return nums[endIndex-1]
	} else {
		// 小于3直接排序替换处理即可
		if nums[startIndex] > nums[endIndex] {
			nums[startIndex], nums[endIndex] = nums[endIndex], nums[startIndex]
		}
		return 0
	}
}
func TestQuickSort(t *testing.T) {
	nums := []int{1, 3, 2, 5, 4, 7, 6}
	QuickSort(nums)
	fmt.Println(nums)
	nums2 := []int{1}
	QuickSort(nums2)
	fmt.Println(nums2)
	nums3 := []int{2, 1}
	QuickSort(nums3)
	fmt.Println(nums3)
	nums4 := []int{1, 3}
	QuickSort(nums4)
	fmt.Println(nums4)
}
