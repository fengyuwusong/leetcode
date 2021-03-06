package easy

import (
	"fmt"
	"testing"
)

/**
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	var lastValue, nowValue int
	for n := 1; n <= rowIndex+1; n++ {
		if n == 1 {
			ans[0] = 1
		} else {
			ans[n-1] = 1
			lastValue, nowValue = ans[0], ans[1]
			for i := 1; i < n-1; i++ {
				nowValue = ans[i]
				ans[i] = lastValue + ans[i]
				lastValue = nowValue
			}
		}
	}
	return ans
}

func TestGetRow(t *testing.T) {
	fmt.Println(getRow(3))
}
