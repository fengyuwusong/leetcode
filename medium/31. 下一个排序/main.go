package main

import "fmt"

// 数学规律解法
/**
1. 首先从后向前查找第一个顺序对 (i,i+1)，满足 a[i] < a[i+1]。这样「较小数」即为 a[i]。此时 [i+1,n) 必然是下降序列。
2. 如果找到了顺序对，那么在区间 [i+1,n)[i+1,n) 中从后向前查找第一个元素 jj 满足 a[i] < a[j]。这样「较大数」即为 a[j]
3. 交换 a[i]与 a[j]，此时可以证明区间 [i+1,n) 必为降序。我们可以直接使用双指针反转区间 [i+1,n) 使其变为升序，而无需对该区间进行排序。
*/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 找到第一个 nums[i]< nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 找到一个nums[j]大于nums[i] 替换nums[i]与nums[j]
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 如果i<0 说明整个数组降序 直接翻转整个数组即可
	// 反转[i+1,n)
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
