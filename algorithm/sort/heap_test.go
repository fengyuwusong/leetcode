package sort

import (
	"fmt"
	"testing"
)

// 堆排序
// 最坏、最好、平均 时间复杂度O(nlogn) 空间复杂度O(logn)
// 不稳定排序
// 思路: https://www.cnblogs.com/chengxiao/p/6129630.html
// 堆是具有以下性质的 完全二叉树 :
// 每个结点的值都大于或等于其左右孩子结点的值，称为大顶堆；
// 或者每个结点的值都小于或等于其左右孩子结点的值，称为小顶堆。
// 将大顶堆按层级进行编号映射到数组中为
// 0   1  2   3   4   5   6   7   8
// 50, 45 40, 20, 25, 35, 30, 10, 15
// 小顶堆
// 0   1   2   3   4   5   6   7   8
// 10, 20, 15, 25, 50, 30, 40, 35, 45

// 即数组具有以下规律
// 大顶堆 arr[i] >= arr[2i+1] && arr[i]>= arr[2i+2] => 父节点大于等于左右子节点
// 即 50>=45&&50>=40 | 45>=20&45>=25
// 小顶堆反之

// 堆排序的基本思想:
// 将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。
// 将其与末尾元素进行交换，此时末尾就为最大值。
// 然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。
// 如此反复执行，便能得到一个有序序列
// 分别有以下步骤
// a.将无需序列构建成一个堆，根据升序降序需求选择大顶堆或小顶堆;
// b.将堆顶元素与末尾元素交换，将最大元素"沉"到数组末端;
// c.重新调整结构，使其满足堆定义，然后继续交换堆顶元素与当前末尾元素，反复执行调整+交换步骤，直到整个序列有序。

func heapSort(nums []int) {
	// 1. 构建大顶堆
	// 从第一个非叶子节点len/2 -1 从下至上 从右至左调整结构
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}

	// 调整堆结构 交换堆顶元素与末尾元素
	for j := len(nums) - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		// 重新对堆进行调整
		adjustHeap(nums, 0, j)
	}
}

func adjustHeap(nums []int, i, numsLen int) {
	// 父节点值
	temp := nums[i]
	// 从i节点的左子节点开始 即 2i+1开始
	for k := i*2 + 1; k < numsLen; k = k*2 + 1 {
		// 左子节点小于右子节点 k指向右子节点
		if k+1 < numsLen && nums[k] < nums[k+1] {
			k++
		}
		// 子节点大于父节点 将子节点赋值父节点 则需要调整的节点变为子节点k i=k 下次遍历则从k的子节点 k=k*2+1 开始
		if nums[k] > temp {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}
	nums[i] = temp
}

func TestHeapSort(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heapSort(nums)
	fmt.Println(nums)
	nums2 := []int{1, 3, 2, 6, 2}
	heapSort(nums2)
	fmt.Println(nums2)
}
