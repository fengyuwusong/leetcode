package medium

import (
	"fmt"
	"testing"
)

/**
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:

输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 暴力 时间复杂度O(n^4) 空间复杂度O(1)
// 解法2 hash 需要考虑hashMap存储的key-value是什么
// 因为A/B/C/D 内数组元素并不是唯一的, 故value并不能仅单存记录是否存在 而是应该记录该值出现的次数
// key有多种选择
// 1. 通过三层遍历记录 key = A[i]+B[j]+C[k] 的值, 那么则仅需在遍历D查找-key的元素即可
// 时间复杂度O(n^3+n)=O(n^3) 空间复杂度O(3n)=O(n)
// 2. 反过来记录 key=A[i] 那么需进行三层遍历计算 B[i]+C[j]+D[k] = -key 时间空间复杂度同上
// 3. 如通过两层遍历记录 key = A[i]+B[j] 那么同样仅需要两层遍历计算 C[i]+D[j] = -key
// 时间复杂度O(n^2) 空间复杂度O(2n)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	// 提前设定容量避免hashMap浪费时间扩容增加时间复杂度
	hashMap := make(map[int]int, len(A)+len(B))
	var sum int
	for _, a := range A {
		for _, b := range B {
			hashMap[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			if num, ok := hashMap[-c-d]; ok {
				sum += num
			}
		}
	}
	return sum
}

func TestFourSumCount(t *testing.T) {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
