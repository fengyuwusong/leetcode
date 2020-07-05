package medium

import (
	"fmt"
	"testing"
)

/**
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。



示例：

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。


提示：

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 暴力 时间复杂度O(n*m^2)=>O(n^3) 空间复杂度O(1)
// 解法2 观察暴力解法 发现具有重复计算过程 可使用dp优化 时间复杂度O(n*m) 空间复杂度O(n*m)
// dp[i][j]表示A[i:]和B[j:]的最长公共前缀 dp[i][j] = dp[i+1][j+1] + 1 (A[i]=B[j]) 否则dp[i][j]=0
// 也可表示为dp[i][j] 表示A[:i]和B[:j]的最长公共前缀 则dp[i][j]=dp[i-1][j-1]+1 (A[i]=B[j]) 否则dp[i][j]=0
// 解法3 滑动窗口 遍历对齐重复子数组 后再对重复子数组进行遍历计算长度
// 时间复杂度 O((n+m)*min(n,m)) (n+m)两个数组分别移动遍历对齐 min(n,m)即最多只需比较最短长度数组即可 空间复杂度O(1)
// 解法4 二分查找+哈希 详情参考
// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/zui-chang-zhong-fu-zi-shu-zu-by-leetcode-solution/
// 即将前缀使用 Rabin-Karp 算法进行hash 那么可通过逆算法的方式使用二分法计算k值最长公共子数组长度

// 解法2
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	// 初始化dp数组
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	// 遍历A数组
	for i := n - 1; i >= 0; i-- {
		// 遍历B数组
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func TestFindLength(t *testing.T) {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
