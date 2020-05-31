package hard

import (
	"fmt"
	"math"
	"testing"
)

// 解法1: 归并两个数组 取得中位数 时间复杂度O(n + m) 空间复杂度O(n+m)

// 解法2: 双指针遍历 当 i + j  = 中位数位置时 即可得到 需考虑各种特殊情况 如 nums1 整体小于 nums2 等
// 时间复杂度O(n+m) 空间复杂度O(1)

// 解法3: 二分查找 O(log(n+m)) 查找第k大的元素 => 分别从两个数组的中间开始查找并排除k/2个元素 直至找到中位数位置
// k = (n+m)/2 或 (n+m)/2+1 那么在A B两个数组的中位数分别为 Ka=A[k/2-1] Kb=B[k/2-1]
// A[k/2-1]前面有 k/2-2个数比A[k/2-1]小, B数组同理,那么可以通过排除缩小k的大小,找到第k大的数。
// 1. 当 A[k/2-1] < B[k/2-1] 时, 可以排除A数组[0, k/2-1]个数, 此时 k = k - k/2 即排除了k/2个数
// 2. 反之当 B[k/2-1] > A[k/2-1]时, 同上.
// 3. 当 A[k/2-1] = B[k/2-1]时，同理, 此时排除A或B数组都可。
// 循环出口 当k=1(AB数组中第一大的数)时,说明仅需排除A[k/2-1] 或 B[k/2-1], 此时取其中最小值即可。
// 按照上述推论求出 偶数: k = (n+m)/2  奇数: k = (n+m)/2, (n+m)/2+1 即可。

// 解法4: 划分数组 时间复杂度O(m+n) 空间复杂度O(1)
//  中位数可以从满足以下两个条件的子集a,b中得到
//          left_part          |         right_part
//    A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]
//    B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]

// 1. a的元素个数>= b的元素个数 (满足第k个概念)
// 2. a的任一元素都小于b的任一元素 (满足排序概念) => 左边数组子集最大值必小于另外一个数组子集的最小值 (本数组已排序,必满足大小关系)
// => B[j−1]<=A[i] 且 A[i−1]<=B[j]

// 那么分析可知 两数组之和奇数时: median = max(A[i-1], b[j-1])
// 偶数时: median = (max(A[i-1], b[j-1]) + min(A[i], B[j])) /2
// => 左边最大值和右边最小值
// 依据上述子集规则可推导:
// 1. i、j之间的关系
// 偶数时条件1: i-1+j-1= m-1+n-1-(i-1+j-1) => i+j = (m+n)/2
// 奇数时条件1: i-1+j-1= m-1+n-1-(i-1+j-1)+1 (+1代表左边比右边大) => i+j = (m+n+1)/2
// 由于在整数除条件下两者相等 (m+n)/2 = (m+n+1)/2 故 i+j = (m+n+1)/2
// 当 A长度必小于B时, 取任意i, 必有 j= (m+n+1)/2 - i

// 2. B[j−1]<=A[i] 且 A[i−1]<=B[j] 推导
// 等i在[0, m]中最大 且 A[i−1]<=B[j] 时, B[j−1]<=A[i] 必定成立
// 2.1 由于 i 从 [0, m]递增,  A[i-1]递增，故必定存在 最大的i满足A[i−1]<=B[j]
// 2.2 由于 i 是最大满足, 故 i+1 不满足, 代入公式 => A[i] > B[j-1] (此处由于i+1, 故j必须-1) 满足 B[j−1]<=A[i]关系

// 由以上两点，可将B[j−1]<=A[i] 且 A[i−1]<=B[j]等价替换为
// 在 [0, m]中，找出最大i值满足 A[i] > B[j-1] ( j=(m+n+1)/2 - i ) 故使用二分查找即可
// 特殊情况考虑:
// 由于需满足 B[j−1]<=A[i] 且 A[i−1]<=B[j] 故 j-1(i-1) 或 i(j) 不能等于 0 或 n/m
// 当i=0时 A[i]=-∞ 当i=n时 i=+∞ j同理 这样两种情况则不会对判断造成影响 应为无论正负无穷对于对于等式来说是不成立的
// 从直观上理解，当一个数组不出现在数组前部分里，对应的值负无穷，必定不会对抢一部分的最大值产生影响, 反之亦然
// 此方法仅需二分遍历AB中最小的数组即可, 故时间复杂度O(log(min(m, n))) 空间复杂度O(1)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len, nums2Len := len(nums1), len(nums2)
	// 确保数组1长度小于数组2 则i必定存在对应的j值
	if nums1Len > nums2Len {
		return findMedianSortedArrays(nums2, nums1)
	}
	// 定义边界
	left, right := 0, nums1Len
	// median = (max(A[i-1], b[j-1]) + min(A[i], B[j])) /2 定义两个数组中间两位数 即左边最大值和右边最小值
	median1, median2 := 0, 0
	// 定义 i j子集分界
	i, j := 0, 0
	// 开始二分遍历
	for left <= right {
		// 从中间值取i
		i = (left + right) / 2
		// j=(m+n+1)/2 - i
		j = (nums1Len+nums2Len+1)/2 - i
		// 排除特定条件
		// i = 0 时 A左边界取无无穷
		nums1LeftNum := math.MinInt64
		if i != 0 {
			nums1LeftNum = nums1[i-1]
		}
		// j同理
		nums2LeftNum := math.MinInt64
		if j != 0 {
			nums2LeftNum = nums2[j-1]
		}
		// i = n 时 取正无穷
		nums1RightNum := math.MaxInt64
		if i != nums1Len {
			nums1RightNum = nums1[i]
		}
		nums2RightNum := math.MaxInt64
		if j != nums2Len {
			nums2RightNum = nums2[j]
		}
		// 判断A[i-1] <= B[j]是否满足
		if nums1LeftNum <= nums2RightNum {
			// 满足则将排除左边界值 取更大值 取右区间 [i+1, right] 再进行二分查找
			left = i + 1
			median1 = Max(nums1LeftNum, nums2LeftNum)
			median2 = Min(nums1RightNum, nums2RightNum)
		} else {
			// 不满足 => A[i-1] > B[j] 排除右区间 取更小值 取左区间 [left, i-1]
			right = i - 1
		}
	}
	// 推出循环说明二分结束 当前i值为最大且符合等式 数组总长度为偶数 则取中间两数平均值 反之取左边最大值(由于左边子集比右边大)
	if (nums1Len+nums2Len)%2 == 0 {
		return float64(median1+median2) / 2
	}
	return float64(median1)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestFindMedianSortedArrays(t *testing.T) {
	//fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{3}, []int{-2, -1}))
}
