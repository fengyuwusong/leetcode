package easy

import (
	"testing"
)

// 7. 整数反转
//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21
//注意:
//
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
func reverse(x int) int {
	var res int
	for {
		if x == 0 {
			break
		}
		res = res*10 + x%10
		x /= 10
	}
	// 根据补码 最大值为除首位全是1 先取全1后右移首位为0  需转换为int类型 否则后续计算数值时不检验首位
	const INT32_MAX = int(^uint32(0) >> 1)
	// 除首位外全是0 取最大值反码即可
	const INT32_MIN = ^INT32_MAX
	if res < INT32_MIN || res > INT32_MAX {
		return 0
	}
	return res
}

func Test_reverseNumber(t *testing.T) {
	println(reverse(123))
}
