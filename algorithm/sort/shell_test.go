package sort

import (
	"fmt"
	"testing"
)

// 插入排序的分组进化版
func shell(nums []int) []int {
	n := len(nums)
	for gap := n / 2; gap > 0; gap /= 2 {
		// 共gap个组 每组进行插入排序
		for i := 0; i < gap; i++ {
			groupSort(nums, n, i, gap)
		}
	}
	return nums
}

/**
 * 对希尔排序中的单个组进行排序
 *
 * 参数说明:
 *     a -- 待排序的数组
 *     n -- 数组总的长度
 *     i -- 组的起始位置
 *     gap -- 组的步长
 *
 *  组是"从i开始，将相隔gap长度的数都取出"所组成的！
 */
func groupSort(a []int, n, i, gap int) {
	for j := i + gap; j < n; j += gap {
		// 如果a[j] < a[j-gap]，则寻找a[j]位置，并将后面数据的位置都后移。
		if a[j] < a[j-gap] {
			tmp := a[j]
			k := j - gap
			for k >= 0 && a[k] > tmp {
				a[k+gap] = a[k]
				k -= gap
			}
			a[k+gap] = tmp
		}
	}
}

func TestShell(t *testing.T) {
	fmt.Printf("%v\n", shell([]int{3, 2, 5, 1, 7, 4}))
}
