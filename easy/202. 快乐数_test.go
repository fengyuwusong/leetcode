package easy

import (
	"fmt"
	"testing"
)

/**
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。



示例：

输入：19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/happy-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 快乐数求解会出现2种情况
// 1. 最终得到1
// 2. 最终进入循环
// 当n位数大于3时 求下一个n=9^2*n = 91*n>243 当n=4 next=324 n=13 next=1053
// 当n位数等于3时 求下一个n=9^2*3 = 243
// 当n位数小于3时 求下一个n=9^*n <243 例如 n=2 next=162 n=1 next=81
// 推倒出结论：当n位数大于3时, 不停求next会使n位数下降 当n小于3时 位数不一定会下降，可能会上升 如n=2 next=162 那么n=3即是可能产生循环的位数
// 故仅需记录每个next值 判断是否循环即可

// 基于以上推论
// 解法1 hash 记录循环
// 时间复杂度 243内最长的一次循环 + 循环求每位的值logn + 高于243数回归到243的时间成本 logn+ loglogn + logloglogn = O(logn)
// 空间复杂度 O(logn) 每个n的sum值需要存储 n共出现logn次
// 解法2 快慢指针 判定循环 时间复杂度 由于需额外进行一次遍历两个指针相等 即O(2logn) = O(logn) 空间复杂度O(1)
// 解法3 可编写程序将243以下的所有循环数找到, 硬编码排除即可 时间复杂度O(logn) 空间复杂度O(1)
func isHappy(n int) bool {
	hashSet := make(map[int]struct{})
	sum := 0
	for sum != 1 {
		sum = 0
		for n != 0 {
			t := n % 10
			n /= 10
			sum += t * t
		}
		if _, ok := hashSet[sum]; ok {
			return false
		}
		n = sum
		hashSet[sum] = struct{}{}
	}
	return true
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}
