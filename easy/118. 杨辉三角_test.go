package easy

import (
	"fmt"
	"testing"
)

/**
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func generate(numRows int) [][]int {
	var ans [][]int
	for n := 1; n <= numRows; n++ {
		if n == 1 {
			ans = append(ans, []int{1})
		} else {
			rows := make([]int, n)
			rows[0], rows[n-1] = 1, 1
			for i := 1; i < n-1; i++ {
				rows[i] = ans[n-2][i-1] + ans[n-2][i]
			}
			ans = append(ans, rows)
		}
	}
	return ans
}
func TestGenerate(t *testing.T) {
	fmt.Println(generate(5))
}
